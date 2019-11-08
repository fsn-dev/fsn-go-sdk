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

var CommandMakeSwap = cli.Command{
	Name:      "makeswap",
	Usage:     "(offline) build make swap raw transaction",
	ArgsUsage: "<fromAssetID> <fomAmount> <toAssetID> <toAmount> <swapSize>",
	Description: `
build make swap raw transaction`,
	Flags: append([]cli.Flag{
		swapFromStartFlag,
		swapFromEndFlag,
		swapToStartFlag,
		swapToEndFlag,
		swapTargetsFlag,
		descriptionFlag,
	}, commonFlags...),
	Action: makeswap,
}

func makeswap(ctx *cli.Context) error {
	if len(ctx.Args()) != 3 {
		cli.ShowCommandHelpAndExit(ctx, "makeswap", 1)
	}

	fromAssetID_ := ctx.Args().Get(0)
	fromAmount_ := ctx.Args().Get(1)
	toAssetID_ := ctx.Args().Get(2)
	toAmount_ := ctx.Args().Get(3)
	swapSize_ := ctx.Args().Get(4)

	fomeAssetID := clicommon.GetHashFromText("fomeAssetID", fromAssetID_)
	fromAmount := clicommon.GetHexBigIntFromText("fromAmount", fromAmount_)
	toAssetID := clicommon.GetHashFromText("toAssetID", toAssetID_)
	toAmount := clicommon.GetHexBigIntFromText("toAmount", toAmount_)
	swapSize := clicommon.GetBigIntFromText("swapSize", swapSize_)

	fromStartTime := getHexUint64Time(ctx, swapFromStartFlag.Name)
	fromEndTime := getHexUint64Time(ctx, swapFromEndFlag.Name)
	toStartTime := getHexUint64Time(ctx, swapToStartFlag.Name)
	toEndTime := getHexUint64Time(ctx, swapToEndFlag.Name)
	targets := clicommon.GetAddressSlice("swapTargets", ctx.StringSlice(swapTargetsFlag.Name))
	description := ctx.String(descriptionFlag.Name)

	// 1. construct corresponding arguments and options
	baseArgs, signOptions := getBaseArgsAndSignOptions(ctx)
	args := &common.MakeSwapArgs{
		FusionBaseArgs: baseArgs,
		FromAssetID:    fomeAssetID,
		FromStartTime:  fromStartTime,
		FromEndTime:    fromEndTime,
		MinFromAmount:  fromAmount,
		ToAssetID:      toAssetID,
		ToStartTime:    toStartTime,
		ToEndTime:      toEndTime,
		MinToAmount:    toAmount,
		SwapSize:       swapSize,
		Targes:         targets,
		Description:    description,
	}

	// 2. check parameters
	now := getNowTime()
	args.Init(getBigIntFromUint64(now))
	if err := args.ToParam().Check(common.BigMaxUint64, now); err != nil {
		utils.Fatalf("check parameter failed: %v", err)
	}

	// 3. build and/or sign transaction through fsnapi
	tx, err := fsnapi.BuildFSNTx(common.MakeSwapFuncExt, args, signOptions)
	if err != nil {
		utils.Fatalf("create tx error: %v", err)
	}

	return printTx(tx, false)
}
