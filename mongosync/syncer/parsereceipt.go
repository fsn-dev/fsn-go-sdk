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

	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/efsn/core/types"
	"github.com/fsn-dev/fsn-go-sdk/efsn/log"
	"github.com/fsn-dev/fsn-go-sdk/mongosync/mongodb"
)

var (
	topicTokenPurchase   = common.HexToHash("0xcd60aa75dea3072fbc07ae6d7d856b5dc5f4eee88854f5b4abf7b680ef8bc50f")
	topicEthPurchase     = common.HexToHash("0x7f4091b46c33e918a0f3aa42307641d17bb67029427a5369e54b353984238705")
	topicAddLiquidity    = common.HexToHash("0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca")
	topicRemoveLiquidity = common.HexToHash("0x0fbf06c058b90cb038a618f8c2acbf6145f8b3570fd1fa56abb8f0f3f05b36e8")
	topicTransfer        = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	topicApproval        = common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
)

func parseReceipt(mt *mongodb.MgoTransaction, receipt *types.Receipt) {
	if receipt == nil || receipt.Status == 0 {
		return
	}
	for idx, rlog := range receipt.Logs {
		if len(rlog.Topics) == 0 {
			continue
		}

		if rlog.Removed {
			continue
		}

		switch rlog.Topics[0] {
		case topicAddLiquidity:
			addExchangeReceipt(mt, rlog, idx, "AddLiquidity")
		case topicRemoveLiquidity:
			addExchangeReceipt(mt, rlog, idx, "RemoveLiquidity")
		case topicTokenPurchase:
			addExchangeReceipt(mt, rlog, idx, "TokenPurchase")
		case topicEthPurchase:
			addExchangeReceipt(mt, rlog, idx, "EthPurchase")
		case topicTransfer:
			addErc20Receipt(mt, rlog, idx, "Transfer")
		case topicApproval:
			addErc20Receipt(mt, rlog, idx, "Approval")
		}
	}
}

func addExchangeReceipt(mt *mongodb.MgoTransaction, rlog *types.Log, logIdx int, logType string) {
	exchange := strings.ToLower(rlog.Address.String())
	topics := rlog.Topics
	address := common.BytesToAddress(topics[1].Bytes())
	fromAmount := new(big.Int).SetBytes(topics[2].Bytes())
	toAmount := new(big.Int).SetBytes(topics[3].Bytes())

	exReceipt := &mongodb.ExchangeReceipt{
		LogType:         logType,
		LogIndex:        logIdx,
		Exchange:        exchange,
		Address:         strings.ToLower(address.String()),
		TokenFromAmount: fromAmount.String(),
		TokenToAmount:   toAmount.String(),
	}

	mt.ExchangeReceipts = append(mt.ExchangeReceipts, exReceipt)
	log.Info("addExchangeReceipt", "receipt", exReceipt)
}

func addErc20Receipt(mt *mongodb.MgoTransaction, rlog *types.Log, logIdx int, logType string) {
	erc20Address := strings.ToLower(rlog.Address.String())
	topics := rlog.Topics
	from := common.BytesToAddress(topics[1].Bytes())
	to := common.BytesToAddress(topics[2].Bytes())
	value := new(big.Int).SetBytes(rlog.Data)

	erc20Receipt := &mongodb.Erc20Receipt{
		LogType:  logType,
		LogIndex: logIdx,
		Erc20:    erc20Address,
		From:     strings.ToLower(from.String()),
		To:       strings.ToLower(to.String()),
		Value:    value.String(),
	}

	mt.Erc20Receipts = append(mt.Erc20Receipts, erc20Receipt)
	log.Info("addErc20Receipt", "receipt", erc20Receipt)
}
