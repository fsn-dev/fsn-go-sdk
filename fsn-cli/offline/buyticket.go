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

package offline

import (
	"time"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/cmd/utils"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/common"
	"github.com/FusionFoundation/fsn-go-sdk/fsnapi"
	"gopkg.in/urfave/cli.v1"
)

var (
	startFlag = cli.Uint64Flag{
		Name:  "start",
		Usage: "ticket start time, 0 means now",
		Value: 0,
	}
	endFlag = cli.Uint64Flag{
		Name:  "end",
		Usage: "ticket end time, 0 means start + 1 month",
		Value: 0,
	}
)

var CommandBuyTicket = cli.Command{
	Name:      "buyticket",
	Usage:     "(offline) build buy ticket raw transaction",
	ArgsUsage: "",
	Description: `
build buy ticket raw transaction`,
	Flags: append([]cli.Flag{
		startFlag,
		endFlag,
	}, commonFlags...),
	Action: buyticket,
}

func buyticket(ctx *cli.Context) error {
	start := getHexUint64(ctx, startFlag.Name)
	end := getHexUint64(ctx, endFlag.Name)

	// 1. construct corresponding arguments and options
	baseArgs, signOptions := getBaseArgsAndSignOptions(ctx)
	args := &common.BuyTicketArgs{
		FusionBaseArgs: baseArgs,
		Start:          start,
		End:            end,
	}

	// 2. check parameters
	now := uint64(time.Now().Unix())
	args.Init(now)
	if err := args.ToParam().Check(common.BigMaxUint64, now, 600); err != nil {
		utils.Fatalf("check parameter failed: %v", err)
	}

	// 3. build and/or sign transaction through fsnapi
	tx, err := fsnapi.BuildFSNTx(common.BuyTicketFunc, args, signOptions)
	if err != nil {
		utils.Fatalf("create tx error: %v", err)
	}

	return printTx(tx, false)
}
