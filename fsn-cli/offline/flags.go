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
	"gopkg.in/urfave/cli.v1"
)

var (
	decodeTxInputFlag = cli.BoolFlag{
		Name:  "input",
		Usage: "decode transaction input data",
	}
	decodeLogDataFlag = cli.BoolFlag{
		Name:  "logdata",
		Usage: "decode transaction receipt log data",
	}
	ticketStartFlag = cli.Uint64Flag{
		Name:  "start",
		Usage: "ticket start time, defaults to now",
	}
	ticketEndFlag = cli.Uint64Flag{
		Name:  "end",
		Usage: "ticket end time, defaults to start + 1 month",
	}
	timeLockStartFlag = cli.Uint64Flag{
		Name:  "start",
		Usage: "time lock start time, defaults to now",
	}
	timeLockEndFlag = cli.Uint64Flag{
		Name:  "end",
		Usage: "time lock end time, defaults to forever",
	}
	nameFlag = cli.StringFlag{
		Name:  "name",
		Usage: "",
	}
	symbolFlag = cli.StringFlag{
		Name:  "symbol",
		Usage: "",
	}
	totalFlag = cli.StringFlag{
		Name:  "total",
		Usage: "",
	}
	decimalsFlag = cli.Uint64Flag{
		Name:  "decimals",
		Usage: "",
		Value: 0,
	}
	canChangeFlag = cli.BoolFlag{
		Name:  "canchange",
		Usage: "can change total supply",
	}
	descriptionFlag = cli.StringFlag{
		Name:  "description",
		Usage: "",
	}
	swapFromAssetIDFlag = cli.StringFlag{
		Name:  "fromasset",
		Usage: "from asset's ID",
	}
	swapFromAmountFlag = cli.StringFlag{
		Name:  "fromamount",
		Usage: "from asset's amount",
	}
	swapFromStartFlag = cli.Uint64Flag{
		Name:  "fromstart",
		Usage: "swap from start time, defaults to now",
	}
	swapFromEndFlag = cli.Uint64Flag{
		Name:  "fromend",
		Usage: "swap from end time, defaults to forever",
	}
	swapToAssetIDFlag = cli.StringFlag{
		Name:  "toasset",
		Usage: "to asset's ID",
	}
	swapToAmountFlag = cli.StringFlag{
		Name:  "toamount",
		Usage: "to asset's amount",
	}
	swapToStartFlag = cli.Uint64Flag{
		Name:  "tostart",
		Usage: "swap to start time, defaults to now",
	}
	swapToEndFlag = cli.Uint64Flag{
		Name:  "toend",
		Usage: "swap to end time, defaults to forever",
	}
	swapTargetsFlag = cli.StringSliceFlag{
		Name:  "target",
		Usage: "private swap addresses",
	}
	swapSwapSizeFlag = cli.StringFlag{
		Name:  "swapsize",
		Usage: "swap size",
	}
	multiSwapFromAssetIDFlag = cli.StringSliceFlag{
		Name:  "fromassets",
		Usage: "from assets's IDs",
	}
	multiSwapFromAmountFlag = cli.StringSliceFlag{
		Name:  "fromamounts",
		Usage: "from assets's amounts",
	}
	multiSwapFromStartFlag = cli.Int64SliceFlag{
		Name:  "fromstarts",
		Usage: "from assets's start times",
	}
	multiSwapFromEndFlag = cli.Int64SliceFlag{
		Name:  "fromends",
		Usage: "from assets's end times",
	}
	multiSwapToAssetIDFlag = cli.StringSliceFlag{
		Name:  "toassets",
		Usage: "to assets's IDs",
	}
	multiSwapToAmountFlag = cli.StringSliceFlag{
		Name:  "toamounts",
		Usage: "to assets's amounts",
	}
	multiSwapToStartFlag = cli.Int64SliceFlag{
		Name:  "tostarts",
		Usage: "to assets's start times",
	}
	multiSwapToEndFlag = cli.Int64SliceFlag{
		Name:  "toends",
		Usage: "to assets's end times",
	}
)
