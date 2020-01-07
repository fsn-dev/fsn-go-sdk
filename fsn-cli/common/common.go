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
	"fmt"
	"math/big"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/cmd/utils"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/common"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/common/hexutil"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/core/types"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/rlp"
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

func GetUint64FromText(whatValue, str string) uint64 {
	bi := GetBigIntFromText(whatValue, str)
	if !bi.IsUint64() {
		utils.Fatalf("Invalid %s value: %s (overflow)", whatValue, str)
	}
	return bi.Uint64()
}

func GetHexUint64(value uint64) *hexutil.Uint64 {
	result := new(hexutil.Uint64)
	*(*uint64)(result) = value
	return result
}

func GetBlockNumberFromText(numStr string) *big.Int {
	number, ok := new(big.Int).SetString(numStr, 0)
	if ok {
		return number
	}
	if numStr != "latest" && numStr != "pending" {
		utils.Fatalf("Invalid block number: '%s'", numStr)
	}
	return nil
}

func GetAddressSlice(whatAddr string, addresses []string) []common.Address {
	result := make([]common.Address, len(addresses))
	for i, addr := range addresses {
		result[i] = GetAddressFromText(whatAddr, addr)
	}
	return result
}

func GetHashSlice(whatHash string, hashes []string) []common.Hash {
	result := make([]common.Hash, len(hashes))
	for i, hash := range hashes {
		result[i] = GetHashFromText(whatHash, hash)
	}
	return result
}

func GetHexBigIntSlice(whatValue string, bigInts []string) []*hexutil.Big {
	result := make([]*hexutil.Big, len(bigInts))
	for i, bi := range bigInts {
		result[i] = GetHexBigIntFromText(whatValue, bi)
	}
	return result
}

func GetHexUint64Slice(whatValue string, nums []int64) []*hexutil.Uint64 {
	result := make([]*hexutil.Uint64, len(nums))
	for i, num := range nums {
		result[i] = GetHexUint64(uint64(num))
	}
	return result
}

func PrintTx(tx *types.Transaction, json bool) error {
	if json {
		bs, err := tx.MarshalJSONWithSender(true)
		if err != nil {
			return fmt.Errorf("json marshal err %v", err)
		}
		fmt.Println(string(bs))
	} else {
		bs, err := rlp.EncodeToBytes(tx)
		if err != nil {
			return fmt.Errorf("rlp encode err %v", err)
		}
		fmt.Println(hexutil.Bytes(bs))
	}
	return nil
}
