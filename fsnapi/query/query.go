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

package query

import (
	"math/big"

	"github.com/fsn-dev/fsn-go-sdk/efsn/ethclient"
)

type QueryClient interface {
	GetAccount(string) (*big.Int, error)
	GetAccountAtBlockNumber(string, int64) (*big.Int, error)
}

type client struct {
	*ethclient.Client
}

func NewClient(url string) (QueryClient, error) {
	ethclient, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return client{ethclient}, nil
}
