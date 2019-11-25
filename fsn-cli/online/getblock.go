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
	"context"
	"fmt"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/common"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/core/types"
	clicommon "github.com/FusionFoundation/fsn-go-sdk/fsn-cli/common"
	"gopkg.in/urfave/cli.v1"
)

var CommandGetBlock = cli.Command{
	Name:      "getblock",
	Category:  "online",
	Usage:     "get block by hash or number",
	ArgsUsage: "<hash_or_number>",
	Description: `
get block by hash or number`,
	Flags: []cli.Flag{
		serverAddrFlag,
	},
	Action: getblock,
}

func getblock(ctx *cli.Context) error {
	setLogger(ctx)
	if len(ctx.Args()) != 1 {
		cli.ShowCommandHelpAndExit(ctx, "getblock", 1)
	}

	client := dialServer(ctx)
	defer client.Close()

	param := ctx.Args().First()

	var block *types.Block
	var hash common.Hash

	err := hash.UnmarshalText([]byte(param))
	if err == nil {
		block, err = client.BlockByHash(context.Background(), hash)
	} else {
		number := clicommon.GetBlockNumberFromText(param)
		block, err = client.BlockByNumber(context.Background(), number)
	}

	if err != nil {
		return err
	}

	bs, err := block.MarshalJSON(true)
	if err != nil {
		return fmt.Errorf("json marshal err %v", err)
	}
	fmt.Println(string(bs))
	return nil
}
