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

package common

import (
	"math/big"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/cmd/utils"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/common"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/common/hexutil"
)

func GetHashFromText(whatHash, hashStr string) (hash common.Hash) {
	err := hash.UnmarshalText([]byte(hashStr))
	if err != nil {
		utils.Fatalf("Invalid %s hash: %s", whatHash, hashStr)
	}
	return hash
}

func GetAddressFromText(whatAddr, addrStr string) (address common.Address) {
	if !common.IsHexAddress(addrStr) {
		utils.Fatalf("Invalid %s address: %s", whatAddr, addrStr)
	}
	return common.HexToAddress(addrStr)
}

func GetBigIntFromText(whatValue, bigIntStr string) *big.Int {
	value, ok := new(big.Int).SetString(bigIntStr, 0)
	if !ok {
		utils.Fatalf("Invalid %s value: %s", whatValue, bigIntStr)
	}
	return value
}

func GetHexBigIntFromText(whatValue, bigIntStr string) *hexutil.Big {
	value, ok := new(big.Int).SetString(bigIntStr, 0)
	if !ok {
		utils.Fatalf("Invalid %s value: %s", whatValue, bigIntStr)
	}
	result := new(hexutil.Big)
	*(*big.Int)(result) = *value
	return result
}

func GetBlockNumberFromText(numStr string) *big.Int {
	number, ok := new(big.Int).SetString(numStr, 0)
	if ok {
		return number
	}
	if numStr != "latest" || numStr != "pending" {
		utils.Fatalf("Invalid block number: %s", numStr)
	}
	return nil
}
