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

package extensions

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/efsn/tools"
	clicommon "github.com/fsn-dev/fsn-go-sdk/fsn-cli/common"
	"gopkg.in/urfave/cli.v1"
)

var CommandCmpTimeLockBalance = cli.Command{
	Name:      "cmptimelockbalance",
	Aliases:   []string{"cmptlb"},
	Category:  "online-extensions",
	Usage:     "cmp time lock balance",
	ArgsUsage: "<assetID> <address> <value>",
	Description: `
cmp time lock balance`,
	Flags: []cli.Flag{
		timeLockStartFlag,
		timeLockEndFlag,
		blockHeightFlag,
		serverAddrFlag,
	},
	Action: cmptimelockbalance,
}

func cmptimelockbalance(ctx *cli.Context) error {
	setLogger(ctx)
	if len(ctx.Args()) != 3 {
		cli.ShowCommandHelpAndExit(ctx, "cmptimelockbalance", 1)
	}

	client := dialServer(ctx)
	defer client.Close()

	assetID := clicommon.GetHashFromText("assetID", ctx.Args().Get(0))
	address := clicommon.GetAddressFromText("address", ctx.Args().Get(1))
	value := clicommon.GetBigIntFromText("value", ctx.Args().Get(2))
	start := getUint64Time(ctx, timeLockStartFlag.Name, common.TimeLockNow)
	end := getUint64Time(ctx, timeLockEndFlag.Name, common.TimeLockForever)
	blockNr := clicommon.GetBlockNumberFromText(ctx.String(blockHeightFlag.Name))

	if start == 0 {
		start = uint64(time.Now().Unix())
	}
	if end == 0 {
		end = common.TimeLockForever
	}

	timelock := common.NewTimeLock(&common.TimeLockItem{
		StartTime: start,
		EndTime:   end,
		Value:     value,
	})

	if err := timelock.IsValid(); err != nil {
		return err
	}

	balance, err := client.GetRawTimeLockBalance(context.Background(), assetID, address, blockNr)
	if err != nil {
		return err
	}

	cmp := balance.Cmp(timelock)
	if cmp >= 0 {
		fmt.Println("Has enough timelock balance of")
		tools.MustPrintJSON(timelock)
		return nil
	}

	diff := &common.TimeLock{}
	x := start
	for _, item := range balance.Items {
		if x > end {
			break
		}
		if item.EndTime < start {
			continue
		}
		if item.StartTime > x { // hole
			diff.Items = append(diff.Items, &common.TimeLockItem{
				StartTime: x,
				EndTime:   common.MinUint64(item.StartTime-1, end),
				Value:     value,
			})
		}
		if item.EndTime != common.TimeLockForever {
			x = item.EndTime + 1
		} else {
			x = item.EndTime
		}
		if item.Value.Cmp(value) >= 0 {
			continue
		}
		if item.StartTime > end {
			break
		}
		// intersection
		diff.Items = append(diff.Items, &common.TimeLockItem{
			StartTime: common.MaxUint64(item.StartTime, start),
			EndTime:   common.MinUint64(item.EndTime, end),
			Value:     new(big.Int).Sub(value, item.Value),
		})
	}
	if x < end { // last hole
		diff.Items = append(diff.Items, &common.TimeLockItem{
			StartTime: x,
			EndTime:   end,
			Value:     value,
		})
	}

	fmt.Println("Not enough timelock balance, differnce is")
	tools.MustPrintJSON(diff)
	return nil
}
