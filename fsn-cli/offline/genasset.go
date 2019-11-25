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
	"github.com/FusionFoundation/fsn-go-sdk/fsnapi"
	"gopkg.in/urfave/cli.v1"
)

var CommandGenAsset = cli.Command{
	Name:      "genasset",
	Category:  "offline",
	Usage:     "build generate asset raw transaction",
	ArgsUsage: "",
	Description: `
build generate asset raw transaction`,
	Flags: append([]cli.Flag{
		nameFlag,
		symbolFlag,
		totalFlag,
		decimalsFlag,
		canChangeFlag,
		descriptionFlag,
	}, commonFlags...),
	Action: genasset,
}

func genasset(ctx *cli.Context) error {
	setLogger(ctx)
	name := ctx.String(nameFlag.Name)
	symbol := ctx.String(symbolFlag.Name)
	total := getHexBigInt(ctx, totalFlag.Name)
	decimals := ctx.Uint64(decimalsFlag.Name)
	canChange := ctx.Bool(canChangeFlag.Name)
	description := ctx.String(descriptionFlag.Name)

	// 1. construct corresponding arguments and options
	baseArgs, signOptions := getBaseArgsAndSignOptions(ctx)
	args := &common.GenAssetArgs{
		FusionBaseArgs: baseArgs,
		Name:           name,
		Symbol:         symbol,
		Decimals:       uint8(decimals),
		Total:          total,
		CanChange:      canChange,
		Description:    description,
	}

	// 2. check parameters
	if err := args.ToParam().Check(common.BigMaxUint64); err != nil {
		utils.Fatalf("check parameter failed: %v", err)
	}

	// 3. build and/or sign transaction through fsnapi
	tx, err := fsnapi.BuildFSNTx(common.GenAssetFunc, args, signOptions)
	if err != nil {
		utils.Fatalf("create tx error: %v", err)
	}

	return printTx(tx, false)
}
