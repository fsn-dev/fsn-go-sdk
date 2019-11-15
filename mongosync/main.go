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
	"gopkg.in/urfave/cli.v1"
)

var appVersion = "1.0.0"

var app *cli.App

func init() {
	app = utils.NewApp(appVersion, "sync blocks into mongodb")
	app.Commands = []cli.Command{
		commandSyncAll,
	}
	app.Flags = append(app.Flags,
		jobsFlag,
		intervalFlag,
		utils.VerbosityFlag,
		utils.JsonFlag,
	)
}

var (
	jobsFlag = cli.Uint64Flag{
		Name:  "jobs",
		Usage: "number of jobs (1-1000)",
		Value: 10,
	}
	intervalFlag = cli.Uint64Flag{
		Name:  "interval",
		Usage: "interval of blocks to show sync progress",
		Value: 100,
	}
)

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
