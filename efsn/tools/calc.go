package tools

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"math/big"

	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/efsn/core/types"
	"github.com/fsn-dev/fsn-go-sdk/efsn/rlp"
)

func CalcRewards(height *big.Int) *big.Int {
	var i int64
	div2 := big.NewInt(2)
	// initial reward 2.5
	var reward = new(big.Int).Mul(big.NewInt(25), big.NewInt(100000000000000000))
	// every 4915200 blocks divide reward by 2
	segment := new(big.Int).Div(height, new(big.Int).SetUint64(4915200))
	for i = 0; i < segment.Int64(); i++ {
		reward = new(big.Int).Div(reward, div2)
	}
	return reward
}

func GetBlockReward(block *types.Block, receipts types.Receipts) *big.Int {
	// block creation reward
	reward := CalcRewards(block.Number())
	gasUses := make(map[common.Hash]uint64)
	for _, receipt := range receipts {
		if receipt != nil {
			gasUses[receipt.TxHash] = receipt.GasUsed
		}
	}
	for _, tx := range block.Transactions() {
		if gasUsed, ok := gasUses[tx.Hash()]; ok {
			gasReward := new(big.Int).Mul(tx.GasPrice(), new(big.Int).SetUint64(gasUsed))
			if gasReward.Sign() > 0 {
				// transaction gas reward
				reward.Add(reward, gasReward)
			}
		}
		if common.IsFsnCall(tx.To()) {
			fsnCallParam := &common.FSNCallParam{}
			rlp.DecodeBytes(tx.Data(), fsnCallParam)
			feeReward := common.GetFsnCallFee(tx.To(), fsnCallParam.Func)
			if feeReward.Sign() > 0 {
				// transaction fee reward
				reward.Add(reward, feeReward)
			}
		}
	}
	return reward
}

func DecodeAllTickets(data []byte) (common.TicketsDataSlice, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read tickets zip data: %v", err)
	}
	var buf bytes.Buffer
	if _, err = io.Copy(&buf, gz); err != nil {
		return nil, fmt.Errorf("Copy tickets zip data: %v", err)
	}
	if err := gz.Close(); err != nil {
		return nil, fmt.Errorf("Close read zip tickets: %v", err)
	}

	var tickets common.TicketsDataSlice
	if err := rlp.DecodeBytes(buf.Bytes(), &tickets); err != nil {
		return nil, fmt.Errorf("Unable to decode tickets, err: %v", err)
	}
	return tickets, nil
}
