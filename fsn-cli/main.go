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

package main

import (
	"fmt"
	"os"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/cmd/utils"
	"github.com/FusionFoundation/fsn-go-sdk/fsn-cli/offline"
	"github.com/FusionFoundation/fsn-go-sdk/fsn-cli/online"
	"gopkg.in/urfave/cli.v1"
)

var appVersion = "0.1.0"

var app *cli.App

func init() {
	app = utils.NewApp(appVersion, "Fusion blockchain client")
	app.Commands = []cli.Command{
		// offline commands
		offline.CommandDecodeRawTx,
		offline.CommandGenAsset,
		offline.CommandSendAsset,
		offline.CommandDecAsset,
		offline.CommandIncAsset,
		offline.CommandBuyTicket,
		offline.CommandGenNotation,
		offline.CommandAsset2TimeLock,
		offline.CommandTimeLock2Asset,
		offline.CommandTimeLock2Timelock,
		offline.CommandMakeSwap,
		offline.CommandTakeSwap,
		offline.CommandRecallSwap,
		offline.CommandMakeMultiSwap,
		offline.CommandTakeMultiSwap,
		offline.CommandRecallMultiSwap,
		// online commands
		online.CommandGetBlock,
		online.CommandGetTransaction,
		online.CommandGetTransactionReceipt,
		online.CommandGetTransactionAndReceipt,
		online.CommandGetTransactionCount,
		online.CommandGetSnapshot,
		online.CommandGetSnapshotAtHash,
		online.CommandSendRawTx,
		online.CommandGetAsset,
		online.CommandGetSwap,
		online.CommandGetNotation,
		online.CommandGetLatestNotation,
		online.CommandGetAddressByNotation,
		online.CommandGetBalance,
		online.CommandGetAllBalances,
		online.CommandGetTimeLockBalance,
		online.CommandGetAllTimeLockBalances,
		online.CommandIsAutoBuyTicket,
		online.CommandTicketPrice,
		online.CommandTotalNumberOfTickets,
		online.CommandAllTickets,
		online.CommandAllInfoByAddress,
		online.CommandGetStakeInfo,
		online.CommandGetBlockReward,
	}
	app.Flags = append(app.Flags, utils.VerbosityFlag)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
