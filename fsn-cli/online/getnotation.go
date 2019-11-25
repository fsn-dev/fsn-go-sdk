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

var CommandGetNotation = cli.Command{
	Name:      "getnotation",
	Category:  "online",
	Usage:     "get notation",
	ArgsUsage: "<address>",
	Description: `
get notation by address`,
	Flags: []cli.Flag{
		blockHeightFlag,
		serverAddrFlag,
	},
	Action: getnotation,
}

func getnotation(ctx *cli.Context) error {
	setLogger(ctx)
	if len(ctx.Args()) != 1 {
		cli.ShowCommandHelpAndExit(ctx, "getnotation", 1)
	}

	client := dialServer(ctx)
	defer client.Close()

	address := clicommon.GetAddressFromText("address", ctx.Args().First())
	blockNr := clicommon.GetBlockNumberFromText(ctx.String(blockHeightFlag.Name))
	notation, err := client.GetNotation(context.Background(), address, blockNr)
	if err != nil {
		return err
	}

	tools.MustPrintJSON(notation)
	return nil
}
