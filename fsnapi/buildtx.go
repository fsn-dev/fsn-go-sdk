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
	"math/big"

	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/efsn/core/types"
)

func GenTxInput(funcType common.FSNCallFunc, args common.FSNBaseArgsInterface) ([]byte, error) {
	funcData, err := args.ToData()
	if err != nil {
		return nil, err
	}
	var param = common.FSNCallParam{Func: funcType, Data: funcData}
	input, err := param.ToBytes()
	if err != nil {
		return nil, err
	}
	return input, nil
}

func BuildFSNTx(funcType common.FSNCallFunc, args common.FSNBaseArgsInterface, signOptions *SignOptions) (tx *types.Transaction, err error) {
	input, err := GenTxInput(funcType, args)
	if err != nil {
		return nil, err
	}
	baseArgs := args.BaseArgs()
	var (
		nonce    uint64
		gasLimit uint64
		gasPrice *big.Int
	)
	if baseArgs.Nonce != nil {
		nonce = uint64(*baseArgs.Nonce)
	}
	if baseArgs.Gas != nil {
		gasLimit = uint64(*baseArgs.Gas)
	}
	if baseArgs.GasPrice != nil {
		gasPrice = baseArgs.GasPrice.ToInt()
	}
	tx = types.NewTransaction(
		nonce,
		common.FSNCallAddress,
		nil,
		gasLimit,
		gasPrice,
		input,
	)
	if signOptions != nil {
		tx, err = SignTx(tx, signOptions)
		if err != nil {
			return nil, err
		}
	}
	return tx, nil
}
