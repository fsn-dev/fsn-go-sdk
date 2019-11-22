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
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/accounts/keystore"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/common"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/core/types"
)

type SignOptions struct {
	Signer   common.Address
	Keyfile  string
	Passfile string
	Password string
	ChainID  uint64
}

func SignTx(tx *types.Transaction, signOptions *SignOptions) (*types.Transaction, error) {
	keyjson, err := ioutil.ReadFile(signOptions.Keyfile)
	if err != nil {
		return nil, err
	}

	passphrase := signOptions.Password
	if passphrase == "" && signOptions.Passfile != "" {
		passdata, err := ioutil.ReadFile(signOptions.Passfile)
		if err != nil {
			return nil, err
		}
		passphrase = strings.TrimSpace(string(passdata))
	}

	key, err := keystore.DecryptKey(keyjson, passphrase)
	if err != nil {
		return nil, err
	}

	if key.Address != signOptions.Signer {
		return nil, fmt.Errorf("key content mismatch: have account %x, want %x", key.Address, signOptions.Signer)
	}

	chainID := new(big.Int).SetUint64(signOptions.ChainID)
	signer := types.NewEIP155Signer(chainID)
	return types.SignTx(tx, signer, key.PrivateKey)
}
