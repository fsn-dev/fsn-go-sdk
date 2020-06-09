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
	"math/big"
	"strings"
	"sync"

	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common/hexutil"
	"github.com/fsn-dev/fsn-go-sdk/efsn/core/types"
	"github.com/fsn-dev/fsn-go-sdk/efsn/log"
	"github.com/fsn-dev/fsn-go-sdk/efsn/rlp"
	"github.com/fsn-dev/fsn-go-sdk/efsn/tools"
	"github.com/fsn-dev/fsn-go-sdk/mongosync/mongodb"
	"gopkg.in/mgo.v2"
)

const (
	averageBlockTime uint64 = 13
	defaultCoinType  string = "FSN"
	defaultTxType    string = "Origin"
	defaultTryTimes  int    = 3
	maxParseBlocks   int    = 1000
)

func tryDoTimes(name string, f func() error) {
	var err error
	i := 0
	for ; i < defaultTryTimes; i++ {
		err = f()
		if err == nil || mgo.IsDup(err) {
			return
		}
	}
	log.Warn("tryDoTimes", "name", name, "time", i, "err", err)
}

func (w *Worker) Parse(block *types.Block, receipts types.Receipts) {
	msg := &Message{
		block:    block,
		receipts: receipts,
	}
	w.messageChan <- msg
}

func (w *Worker) startParser(wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	wg2 := new(sync.WaitGroup)
	for {
		select {
		case msg := <-w.messageChan:
			if msg == nil {
				return
			}
			count++
			wg2.Add(2)
			// parse block
			go w.parseBlock(msg.block, msg.receipts, wg2)
			// parse transactions
			go w.parseTransactions(msg.block, msg.receipts, wg2)
			if count == maxParseBlocks {
				count = 0
				wg2.Wait() // prevent memory exhausted (when blocks too large)

			}
		}
	}
	wg2.Wait()
}

func (w *Worker) parseBlock(block *types.Block, receipts types.Receipts, wg *sync.WaitGroup) {
	defer wg.Done()
	mb := new(mongodb.MgoBlock)

	hash := block.Hash().String()

	mb.Key = hash
	mb.Number = block.NumberU64()
	mb.Hash = hash
	mb.ParentHash = block.ParentHash().String()
	mb.Nonce = block.Nonce()
	mb.Miner = strings.ToLower(block.Coinbase().String())
	mb.Difficulty = block.Difficulty().Uint64()
	mb.TotalDifficulty = 0
	mb.Size = uint64(block.Size())
	mb.GasLimit = block.GasLimit()
	mb.GasUsed = block.GasUsed()
	mb.Timestamp = block.Time().Uint64()
	mb.BlockTime = 0
	mb.TxCount = block.Transactions().Len()
	mb.AvgGasprice = calcAvgGasprice(block.Transactions()).String()
	mb.Reward = calcBlockReward(block, receipts).String()

	w.parseSnapshot(mb, block.Header())

	tryDoTimes("AddBlock "+hash, func() error {
		return mongodb.AddBlock(mb, Overwrite)
	})
}

func calcAvgGasprice(txs types.Transactions) *big.Int {
	txCount := int64(len(txs))
	if txCount == 0 {
		return big.NewInt(0)
	}
	sum := big.NewInt(0)
	for _, tx := range txs {
		sum = sum.Add(sum, tx.GasPrice())
	}
	return new(big.Int).Div(sum, big.NewInt(txCount))
}

func calcBlockReward(block *types.Block, receipts types.Receipts) *big.Int {
	reward := big.NewInt(0)
	if block.Number().Sign() > 0 {
		reward = tools.GetBlockReward(block, receipts)
	}
	return reward
}

func (w *Worker) parseSnapshot(mb *mongodb.MgoBlock, header *types.Header) {
	// Fusion related
	snap, err := tools.NewSnapshotFromHeader(header)
	if err != nil {
		return
	}
	mb.SelectedTicket = snap.Selected.String()
	mb.TicketNumber = snap.TicketNumber
	retreatCount := len(snap.Retreat)
	if retreatCount > 0 && header.Number.Sign() > 0 {
		prevHeight := new(big.Int).Sub(header.Number, big.NewInt(1))
		owners := w.getTicketsOwner(snap.Retreat, prevHeight)
		mb.RetreatTickets = make([]string, retreatCount)
		mb.RetreatMiners = make([]string, retreatCount)
		for i := 0; i < retreatCount; i++ {
			mb.RetreatTickets[i] = snap.Retreat[i].String()
			mb.RetreatMiners[i] = strings.ToLower(owners[i].String())
		}
	}
}

func (w *Worker) parseTransactions(block *types.Block, receipts types.Receipts, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(len(block.Transactions()))
	for i, tx := range block.Transactions() {
		go w.parseTx(i, tx, block, receipts, wg)
	}
}

func (w *Worker) parseTx(i int, tx *types.Transaction, block *types.Block, receipts types.Receipts, wg *sync.WaitGroup) {
	defer wg.Done()
	mt := new(mongodb.MgoTransaction)

	receipt := receipts[i]
	hash := tx.Hash().String()

	mt.Key = hash
	mt.Hash = hash
	mt.Nonce = tx.Nonce()
	mt.BlockHash = block.Hash().String()
	mt.BlockNumber = block.NumberU64()
	mt.TransactionIndex = i
	mt.From = strings.ToLower(getTxSender(tx).String())
	mt.To = "nil"
	if tx.To() != nil {
		mt.To = strings.ToLower(tx.To().String())
	}
	txValue := tx.Value()
	mt.Value = txValue.String()
	if txValue.Sign() > 0 {
		mt.ValueInt = new(big.Int).Div(txValue, big.NewInt(1e18)).Uint64()
		mt.ValueDec = new(big.Int).Mod(txValue, big.NewInt(1e18)).Uint64()
	}
	mt.GasLimit = tx.Gas()
	mt.GasPrice = tx.GasPrice().String()
	if receipt != nil {
		mt.GasUsed = receipt.GasUsed
		mt.Status = receipt.Status
	}
	mt.Timestamp = block.Time().Uint64()
	txData := tx.Data()
	if len(txData) > 0 {
		mt.Input = hexutil.Encode(txData)
	}
	mt.CoinType = defaultCoinType
	mt.Type = defaultTxType

	// Fusion related
	if common.IsFsnCall(tx.To()) {
		parseFsnTx(mt, tx, receipt)
	} else if receipt != nil && len(receipt.Logs) != 0 {
		mt.Log = parseReceiptLogs(receipt.Logs) // mt.Log
	}

	tryDoTimes("AddTransaction "+hash, func() error {
		return mongodb.AddTransaction(mt, Overwrite)
	})
}

func hashesToStrings(hashes []common.Hash) []string {
	res := make([]string, len(hashes))
	for k, v := range hashes {
		res[k] = v.String()
	}
	return res
}

func parseReceiptLogs(rlogs []*types.Log) []map[string]interface{} {
	if len(rlogs) == 0 {
		return nil
	}
	logs := make([]map[string]interface{}, len(rlogs))
	for k, v := range rlogs {
		logs[k] = make(map[string]interface{})
		res := logs[k]
		res["contract"] = strings.ToLower(v.Address.String())
		if len(v.Topics) == 0 {
			continue
		}
		switch v.Topics[0] {
		case common.LogFusionAssetReceivedTopic, common.LogFusionAssetSentTopic:
			isReceive := v.Topics[0] == common.LogFusionAssetReceivedTopic
			res["asset"] = v.Topics[1].String()
			addr := strings.ToLower(common.BytesToAddress(v.Topics[2].Bytes()).String())
			if isReceive {
				res["from"] = addr
				res["topic"] = "TimeLockContractReceive"
			} else {
				res["to"] = addr
				res["topic"] = "TimeLockContractSend"
			}
			res["value"] = common.GetBigInt(v.Data, 0, 32).String()
			res["start"] = common.GetBigInt(v.Data, 32, 32).String()
			res["end"] = common.GetBigInt(v.Data, 64, 32).String()
			res["flag"] = common.GetBigInt(v.Data, 96, 32).String()
		default:
			res["topics"] = hashesToStrings(v.Topics)
			res["data"] = hexutil.Encode(v.Data)
		}
	}
	return logs
}

func parseFsnTx(mt *mongodb.MgoTransaction, tx *types.Transaction, receipt *types.Receipt) {
	var fsnCall common.FSNCallParam
	rlp.DecodeBytes(tx.Data(), &fsnCall)
	mt.Type = fsnCall.Func.Name() // mt.Type

	if receipt == nil || len(receipt.Logs) == 0 {
		return
	}
	logData, err := tools.DecodeLogData(receipt.Logs[0].Data)
	if err != nil {
		return
	}
	logMap, ok := logData.(map[string]interface{})
	if !ok {
		return
	}
	if fsnCall.Func == common.ReportIllegalFunc {
		if del, ok := logMap["DeleteTickets"]; ok {
			if delstr, ok := del.(string); ok {
				deltickets, err := tools.DecodePunishTickets(delstr)
				if err == nil {
					logMap["DeleteTickets"] = hashesToStrings(deltickets)
				}
			}
		}
	}
	mt.Log = logMap // mt.Log

	if assetID, ok := logMap["AssetID"].(string); ok {
		if common.HexToHash(assetID) != common.SystemAssetID {
			mt.CoinType = assetID // mt.CoinType
		}
	}

	if _, hasError := logMap["Error"]; hasError {
		mt.Status = types.ReceiptStatusFailed // mt.Status
		return
	}

	if fsnCall.Func == common.SendAssetFunc && mt.CoinType == defaultCoinType {
		sendAssetParam := common.SendAssetParam{}
		rlp.DecodeBytes(fsnCall.Data, &sendAssetParam)
		mt.To = strings.ToLower(sendAssetParam.To.String()) // mt.To
		txValue := sendAssetParam.Value
		mt.Value = txValue.String() // mt.Value
		if txValue.Sign() > 0 {
			mt.ValueInt = new(big.Int).Div(txValue, big.NewInt(1e18)).Uint64() // mt.ValueInt
			mt.ValueDec = new(big.Int).Mod(txValue, big.NewInt(1e18)).Uint64() // mt.ValueDec
		}
	}
}

func getTxSender(tx *types.Transaction) common.Address {
	signer := types.NewEIP155Signer(tx.ChainId())
	sender, _ := types.Sender(signer, tx)
	return sender
}
