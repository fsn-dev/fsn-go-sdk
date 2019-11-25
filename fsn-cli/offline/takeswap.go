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
	"github.com/FusionFoundation/fsn-go-sdk/efsn/cmd/utils"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/common"
	clicommon "github.com/FusionFoundation/fsn-go-sdk/fsn-cli/common"
	"github.com/FusionFoundation/fsn-go-sdk/fsnapi"
	"gopkg.in/urfave/cli.v1"
)

var CommandTakeSwap = cli.Command{
	Name:      "takeswap",
	Category:  "offline",
	Usage:     "build take swap raw transaction",
	ArgsUsage: "<swapID> <size>",
	Description: `
build take swap raw transaction`,
	Flags:  append([]cli.Flag{}, commonFlags...),
	Action: takeswap,
}

func takeswap(ctx *cli.Context) error {
	setLogger(ctx)
	if len(ctx.Args()) != 2 {
		cli.ShowCommandHelpAndExit(ctx, "takeswap", 1)
	}

	swapID_ := ctx.Args().Get(0)
	size_ := ctx.Args().Get(1)

	swapID := clicommon.GetHashFromText("swapID", swapID_)
	size := clicommon.GetBigIntFromText("size", size_)

	// 1. construct corresponding arguments and options
	baseArgs, signOptions := getBaseArgsAndSignOptions(ctx)
	args := &common.TakeSwapArgs{
		FusionBaseArgs: baseArgs,
		SwapID:         swapID,
		Size:           size,
	}

	// 2. check parameters
	if args.Size.Sign() == 0 {
		utils.Fatalf("check parameter failed, wrong size %s", size_)
	}

	// 3. build and/or sign transaction through fsnapi
	tx, err := fsnapi.BuildFSNTx(common.TakeSwapFuncExt, args, signOptions)
	if err != nil {
		utils.Fatalf("create tx error: %v", err)
	}

	return printTx(tx, false)
}
