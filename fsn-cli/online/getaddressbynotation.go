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

	"github.com/FusionFoundation/fsn-go-sdk/efsn/tools"
	clicommon "github.com/FusionFoundation/fsn-go-sdk/fsn-cli/common"
	"gopkg.in/urfave/cli.v1"
)

var CommandGetAddressByNotation = cli.Command{
	Name:      "getaddressbynotation",
	Category:  "online",
	Usage:     "get address by notation",
	ArgsUsage: "<notation>",
	Description: `
get address by notation`,
	Flags: []cli.Flag{
		blockHeightFlag,
		serverAddrFlag,
	},
	Action: getaddressbynotation,
}

func getaddressbynotation(ctx *cli.Context) error {
	setLogger(ctx)
	if len(ctx.Args()) != 1 {
		cli.ShowCommandHelpAndExit(ctx, "getaddressbynotation", 1)
	}

	client := dialServer(ctx)
	defer client.Close()

	notation := clicommon.GetUint64FromText("notation", ctx.Args().First())
	blockNr := clicommon.GetBlockNumberFromText(ctx.String(blockHeightFlag.Name))
	address, err := client.GetAddressByNotation(context.Background(), notation, blockNr)
	if err != nil {
		return err
	}

	tools.MustPrintJSON(address)
	return nil
}
