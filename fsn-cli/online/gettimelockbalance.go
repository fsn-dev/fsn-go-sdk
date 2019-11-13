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

var CommandGetTimeLockBalance = cli.Command{
	Name:      "gettimelockbalance",
	Aliases:   []string{"gettlb"},
	Usage:     "(online) get time lock balance",
	ArgsUsage: "<assetID> <address>",
	Description: `
get time lock balance`,
	Flags: []cli.Flag{
		rawTimeLockFlag,
		blockHeightFlag,
		serverAddrFlag,
	},
	Action: gettimelockbalance,
}

func gettimelockbalance(ctx *cli.Context) error {
	setLogger(ctx)
	if len(ctx.Args()) != 2 {
		cli.ShowCommandHelpAndExit(ctx, "gettimelockbalance", 1)
	}

	client := dialServer(ctx)
	defer client.Close()

	assetID := clicommon.GetHashFromText("assetID", ctx.Args().Get(0))
	address := clicommon.GetAddressFromText("address", ctx.Args().Get(1))
	blockNr := clicommon.GetBlockNumberFromText(ctx.String(blockHeightFlag.Name))

	var balance *interface{}
	var err error
	if ctx.Bool(rawTimeLockFlag.Name) {
		balance, err = client.GetRawTimeLockBalance(context.Background(), assetID, address, blockNr)
	} else {
		balance, err = client.GetTimeLockBalance(context.Background(), assetID, address, blockNr)
	}
	if err != nil {
		return err
	}

	tools.MustPrintJSON(balance)
	return nil
}
