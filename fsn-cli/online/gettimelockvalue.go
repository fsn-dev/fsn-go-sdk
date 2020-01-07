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

var CommandGetTimeLockValue = cli.Command{
	Name:      "getTimeLockValue",
	Aliases:   []string{"gettlv"},
	Category:  "online",
	Usage:     "get time lock value",
	ArgsUsage: "<assetID> <address> <startTime> <endTime>",
	Description: `
get time lock spendable value between specified time range`,
	Flags: []cli.Flag{
		blockHeightFlag,
		serverAddrFlag,
	},
	Action: getTimeLockValue,
}

func getTimeLockValue(ctx *cli.Context) error {
	setLogger(ctx)
	if len(ctx.Args()) != 4 {
		cli.ShowCommandHelpAndExit(ctx, "getTimeLockValue", 1)
	}

	client := dialServer(ctx)
	defer client.Close()

	assetID := clicommon.GetHashFromText("assetID", ctx.Args().Get(0))
	address := clicommon.GetAddressFromText("address", ctx.Args().Get(1))
	startTime := clicommon.GetUint64FromText("startTime", ctx.Args().Get(2))
	endTime := clicommon.GetUint64FromText("endTime", ctx.Args().Get(3))
	blockNr := clicommon.GetBlockNumberFromText(ctx.String(blockHeightFlag.Name))

	balance, err := client.GetTimeLockValue(context.Background(), assetID, address, startTime, endTime, blockNr)
	if err != nil {
		return err
	}

	tools.MustPrintJSON(balance)
	return nil
}
