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
		offline.CommandSendAsset,
		offline.CommandBuyTicket,
		// online commands
		online.CommandSendRawTx,
		online.CommandGetAsset,
	}
	app.Flags = append(app.Flags, utils.VerbosityFlag)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
