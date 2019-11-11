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

var CommandBuyTicket = cli.Command{
	Name:      "buyticket",
	Usage:     "(offline) build buy ticket raw transaction",
	ArgsUsage: "",
	Description: `
build buy ticket raw transaction`,
	Flags: append([]cli.Flag{
		ticketStartFlag,
		ticketEndFlag,
	}, commonFlags...),
	Action: buyticket,
}

func buyticket(ctx *cli.Context) error {
	setLogger(ctx)
	start := getHexUint64Time(ctx, ticketStartFlag.Name)
	end := getHexUint64Time(ctx, ticketEndFlag.Name)

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
