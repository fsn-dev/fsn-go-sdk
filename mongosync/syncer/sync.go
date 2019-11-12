// Copyright 2019 The fsn-go-sdk Authors
// This file is part of the fsn-go-sdk library.
//
// The fsn-go-sdk library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The fsn-go-sdk library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the fsn-go-sdk library. If not, see <http://www.gnu.org/licenses/>.

package syncer

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/common"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/core/types"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/ethclient"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/log"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/tools"
	"github.com/FusionFoundation/fsn-go-sdk/mongosync/mongodb"
)

var (
	ServerURL string
	Overwrite bool
	JobCount  uint64

	minJobs       uint64 = 1
	maxJobs       uint64 = 1000
	minWorkBlocks uint64 = 100

	messageChanSize = 1000

	retryDuration = time.Duration(1) * time.Second
	waitDuration  = time.Duration((averageBlockTime+1)/2) * time.Second

	client     *ethclient.Client
	cliContext = context.Background()
	workers    []*Worker
)

type Message struct {
	block    *types.Block
	receipts types.Receipts
}

type Worker struct {
	id              int // identify worker
	parentTimestamp uint64
	stable          uint64
	start           uint64
	end             uint64

	messageChan chan *Message
}

type Syncer struct {
	stable uint64
	start  uint64
	end    uint64

	last uint64
}

func NewSyncer(stable, start, end uint64) *Syncer {
	return &Syncer{
		stable: stable,
		start:  start,
		end:    end,
	}
}

func SetJobCount(count uint64) {
	if count < minJobs {
		count = minJobs
	} else if count > maxJobs {
		count = maxJobs
	}
	JobCount = count
}

func InitMongoServer(mongoURL, dbName string) {
	mongodb.MongoServerInit(mongoURL, dbName)
}

func DialServer() (err error) {
	client, err = ethclient.Dial(ServerURL)
	if err != nil {
		log.Error("client connection error", "server", ServerURL, "err", err)
		return err
	}
	log.Info("client connection succeed", "server", ServerURL)
	return nil
}

func Close() {
	client.Close()
}

func (s *Syncer) Sync() {
	if err := DialServer(); err != nil {
		return
	}
	defer Close()
	s.dipatchWork()
	s.doWork()
}

func (s *Syncer) dipatchWork() {
	start := s.start
	last := s.end
	if s.start == 0 && s.end == 0 {
		if syncInfo, err := mongodb.FindLatestSyncInfo(); err == nil {
			start = syncInfo.Number
			if start != 0 {
				start++
			}
		}
	}
	for s.end == 0 {
		latestHeader, err := client.HeaderByNumber(cliContext, nil)
		if err == nil {
			last = latestHeader.Number.Uint64()
			if last > s.stable {
				last -= s.stable
			}
			break
		}
		log.Warn("get latest block header failed", "err", err)
		time.Sleep(retryDuration)
	}
	if last <= start && s.end != 0 {
		log.Info("no need to sync block", "begin", start, "end", last)
		return
	}

	s.start = start
	s.last = last

	blockCount := uint64(1)
	if last > start {
		blockCount = last - start
	}
	if blockCount < minWorkBlocks {
		s.last = start
		return
	}
	workerCount := blockCount / minWorkBlocks
	if workerCount > JobCount {
		workerCount = JobCount
	}
	stepCount := blockCount / workerCount

	for i := uint64(0); i < workerCount; i++ {
		wstart := start + i*stepCount
		wend := start + (i+1)*stepCount
		if i == workerCount-1 {
			wend = last
		}
		w := &Worker{
			id:          int(i + 1),
			stable:      s.stable,
			start:       wstart,
			end:         wend,
			messageChan: make(chan *Message, messageChanSize),
		}
		workers = append(workers, w)
	}

	log.Info("dispatch work", "count", workerCount, "step", stepCount, "start", start, "end", last)
}

func (s *Syncer) doWork() {
	mongodb.InitLatestSyncInfo()
	if len(workers) != 0 {
		s.doSyncWork()
	}
	if s.end == 0 {
		s.doLoopWork()
	}
}

func (s *Syncer) fsync() error {
	return mongodb.Fsync(false)
}

func (s *Syncer) checkSync(start, end uint64) {
	s.fsync()
	log.Info("checkSync", "from", start, "to", end)
	checkWorker := &Worker{
		id:          -1,
		stable:      s.stable,
		start:       start,
		end:         end,
		messageChan: make(chan *Message, 10),
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go checkWorker.doSync(wg)
	wg.Wait()
	s.fsync()
}

func (s *Syncer) doSyncWork() {
	log.Info("doSyncWork start", "from", s.start, "to", s.last)
	wg := new(sync.WaitGroup)
	wg.Add(len(workers))
	for _, worker := range workers {
		go worker.doSync(wg)
	}
	wg.Wait()
	log.Info("doSyncWork finished", "from", s.start, "to", s.last)

	log.Info("checkSync start", "from", s.start, "to", s.last)
	s.checkSync(s.start, s.last)
	log.Info("checkSync finished", "from", s.start, "to", s.last)
}

func (s *Syncer) doLoopWork() {
	log.Info("doLoopWork start")
	loopWorker := &Worker{
		id:          0,
		stable:      s.stable,
		start:       s.last,
		end:         0,
		messageChan: make(chan *Message, messageChanSize),
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go loopWorker.doSync(wg)
	go s.UpdateBlockInfo(wg)
	wg.Wait()
	log.Info("doLoopWork finished")
}

func (w *Worker) doSync(wg *sync.WaitGroup) {
	defer func(bstart time.Time) {
		log.Info("End sync process", "id", w.id, "start", w.start, "end", w.end, "duration", common.PrettyDuration(time.Since(bstart)))
		close(w.messageChan)
		wg.Done()
	}(time.Now())

	wg.Add(1)
	go w.startParser(wg)

	log.Info("Start sync process", "id", w.id, "start", w.start, "end", w.end)

	latest := w.end
	height := w.start
	for {
		if w.end > 0 && height >= w.end {
			break
		}
		if height+w.stable > latest {
			latestHeader, err := client.HeaderByNumber(cliContext, nil)
			if err != nil {
				log.Warn("get latest block header failed", "id", w.id, "err", err)
				time.Sleep(retryDuration)
				continue
			}
			latest = latestHeader.Number.Uint64()
			if height+w.stable > latest {
				time.Sleep(waitDuration)
				continue
			}
		}
		last := latest - w.stable
		if w.end > 0 && last >= w.end {
			last = w.end - 1
		}
		w.syncRange(height, last)
		height = last + 1
	}
	w.messageChan <- nil
}

func getSynced(mbs []*mongodb.MgoBlock, num uint64) *mongodb.MgoBlock {
	for _, mb := range mbs {
		if mb.Number == num {
			return mb
		}
	}
	return nil
}

func (w *Worker) syncRange(start, end uint64) {
	step := uint64(10000)
	height := start
	for height <= end {
		from := height
		to := from + step - 1
		if to > end {
			to = end
		}
		mblocks, err := mongodb.FindBlocksInRange(from, to)
		if err != nil {
			log.Error("syncRange error", "from", from, "to", to, "err", err)
			time.Sleep(waitDuration)
			continue
		}
		if len(mblocks) == int(to-from+1) {
			log.Info("syncRange already synced", "id", w.id, "from", from, "to", to)
			height = to + 1
			continue
		}
		if w.end != 0 {
			log.Info("syncRange", "id", w.id, "from", from, "to", to, "exist", len(mblocks))
		}
		for height <= to {
			mb := getSynced(mblocks, height)
			if Overwrite || mb == nil {
				block, err := client.BlockByNumber(cliContext, new(big.Int).SetUint64(height))
				if err != nil {
					log.Warn("get block failed", "id", w.id, "number", height, "err", err)
					time.Sleep(retryDuration)
					continue
				}
				txs := block.Transactions()
				receipts := make(types.Receipts, len(txs))
				wg := new(sync.WaitGroup)
				wg.Add(len(txs))
				for i, tx := range txs {
					go func(index int, txhash common.Hash) {
						defer wg.Done()
						receipt, _ := client.TransactionReceipt(cliContext, txhash)
						receipts[index] = receipt
					}(i, tx.Hash())
				}
				wg.Wait()
				w.Parse(block, receipts)
				if w.end == 0 {
					log.Info("sync block completed", "id", w.id, "number", height)
				} else if height%1000 == 0 {
					log.Info("syncRange in process", "id", w.id, "number", height)
				}
			}
			height += 1
		}
		if w.end != 0 {
			log.Info("syncRange completed", "id", w.id, "from", from, "to", to)
		}
	}
}

func (w *Worker) getTicketsOwner(tids []common.Hash, blockNumber *big.Int) []common.Address {
	owners := make([]common.Address, len(tids))
	if len(tids) == 0 {
		return owners
	}
	data, err := client.CodeAt(cliContext, common.TicketKeyAddress, blockNumber)
	if err != nil {
		return owners
	}
	tickets, err := tools.DecodeAllTickets(data)
	if err != nil {
		return owners
	}
	for i, tid := range tids {
		if ticket, err := tickets.Get(tid); err == nil {
			owners[i] = ticket.Owner
		}
	}
	return owners
}

func (s *Syncer) UpdateBlockInfo(wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		start uint64
		ts    uint64
		td    uint64
	)
	syncInfo, err := mongodb.FindLatestSyncInfo()
	if err == nil && syncInfo.Number != 0 {
		start = syncInfo.Number
		ts = syncInfo.Timestamp
		td = syncInfo.TotalDifficulty
		start++
	}

	last := s.last
	step := uint64(10000)
	height := start
	for {
		from := height
		to := from + step - 1
		if from <= last && to > last {
			to = last
		}
		mblocks, err := mongodb.FindBlocksInRange(from, to)
		if err != nil {
			log.Error("update block error", "from", from, "to", to, "err", err)
			time.Sleep(waitDuration)
			continue
		}
		if to <= last {
			if len(mblocks) != int(to-from+1) {
				s.checkSync(from, to+1)
			}
			log.Info("updateRange", "from", from, "to", to, "exist", len(mblocks))
		}
		for height <= to {
			mb := getSynced(mblocks, height)
			for mb == nil {
				mb, err = mongodb.FindBlockByNumber(height)
				if err == nil {
					break
				}
				if height <= last {
					log.Warn("update block warn", "number", height, "err", "block not synced")
					s.checkSync(height, height+1)
					time.Sleep(10 * waitDuration)
				} else {
					time.Sleep(waitDuration)
				}
			}
			blockTime := mb.Timestamp - ts
			if height <= 1 {
				if height == 0 {
					blockTime = 0
				} else {
					blockTime = averageBlockTime
				}
			}
			totalDifficulty := td + mb.Difficulty
			err = mongodb.UpdateBlockInfo(mb.Key, blockTime, totalDifficulty)
			if err == nil {
				timestamp := mb.Timestamp
				err = mongodb.UpdateSyncInfo(height, timestamp, totalDifficulty)
				if err == nil {
					if height > last {
						log.Info("update block completed", "number", height, "timestamp", timestamp, "td", totalDifficulty)
					} else if height%1000 == 0 {
						log.Info("update block in process", "number", height, "timestamp", timestamp, "td", totalDifficulty)
					}
					td = totalDifficulty
					ts = timestamp
					height++
					continue
				}
			}
			log.Warn("update block error", "number", height, "err", err)
			time.Sleep(waitDuration)
		}
		if to <= last {
			log.Info("updateRange completed", "from", from, "to", to)
		}
		height = to + 1
	}
}
