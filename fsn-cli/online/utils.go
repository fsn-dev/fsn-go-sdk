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

package online

import (
	"github.com/FusionFoundation/fsn-go-sdk/efsn/cmd/utils"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/ethclient"
	"gopkg.in/urfave/cli.v1"
)

var (
	serverAddrFlag = cli.StringFlag{
		Name:  "server",
		Usage: "server address",
		Value: "http://localhost:6509",
	}
	blockHeightFlag = cli.StringFlag{
		Name:  "height",
		Usage: "block height",
		Value: "latest",
	}
)

func dialServer(ctx *cli.Context) *ethclient.Client {
	serverAddr := ctx.String(serverAddrFlag.Name)
	client, err := ethclient.Dial(serverAddr)
	if err != nil {
		utils.Fatalf("dial server %s err %v", serverAddr, err)
	}
	return client
}
