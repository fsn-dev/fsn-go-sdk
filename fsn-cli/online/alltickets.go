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

var CommandAllTickets = cli.Command{
	Name:      "alltickets",
	Aliases:   []string{"gettickets"},
	Category:  "online",
	Usage:     "get tikcets",
	ArgsUsage: "[<address>]",
	Description: `
get all tikcets, or tickets of a specified address`,
	Flags: []cli.Flag{
		blockHeightFlag,
		serverAddrFlag,
	},
	Action: alltickets,
}

func alltickets(ctx *cli.Context) error {
	setLogger(ctx)

	client := dialServer(ctx)
	defer client.Close()

	blockNr := clicommon.GetBlockNumberFromText(ctx.String(blockHeightFlag.Name))
	argsCount := len(ctx.Args())

	if argsCount == 0 {
		tickets, err := client.AllTickets(context.Background(), blockNr)
		if err != nil {
			return err
		}
		tools.MustPrintJSON(tickets)
		return nil
	}

	if argsCount != 1 {
		cli.ShowCommandHelpAndExit(ctx, "alltickets", 1)
	}

	address := clicommon.GetAddressFromText("address", ctx.Args().First())
	tickets, err := client.AllTicketsByAddress(context.Background(), address, blockNr)
	if err != nil {
		return err
	}

	tools.MustPrintJSON(tickets)
	return nil
}
