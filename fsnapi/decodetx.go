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

package fsnapi

import (
	"bytes"
	"fmt"

	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common/hexutil"
	"github.com/fsn-dev/fsn-go-sdk/efsn/core/types"
	"github.com/fsn-dev/fsn-go-sdk/efsn/rlp"
	"github.com/fsn-dev/fsn-go-sdk/efsn/tools"
)

func GetTxSender(tx *types.Transaction) common.Address {
	signer := types.NewEIP155Signer(tx.ChainId())
	sender, _ := types.Sender(signer, tx)
	return sender
}

func DecodeRawTx(hexstr string) (*types.Transaction, error) {
	data, err := hexutil.Decode(hexstr)
	if err != nil {
		return nil, fmt.Errorf("wrong hex string, %v", err)
	}

	var tx types.Transaction
	err = rlp.Decode(bytes.NewReader(data), &tx)
	if err != nil {
		return nil, fmt.Errorf("decode raw transaction err %v", err)
	}
	return &tx, nil
}

func DecodeTxInput(hexstr string) (interface{}, error) {
	data, err := hexutil.Decode(hexstr)
	if err != nil {
		return nil, fmt.Errorf("wrong hex string, %v", err)
	}
	return tools.DecodeTxInput(data)
}

func DecodeTxReceiptLogData(hexstr string) (interface{}, error) {
	data, err := hexutil.Decode(hexstr)
	if err != nil {
		return nil, fmt.Errorf("wrong hex string, %v", err)
	}
	return tools.DecodeLogData(data)
}
