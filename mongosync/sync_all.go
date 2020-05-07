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
	"github.com/FusionFoundation/fsn-go-sdk/efsn/cmd/utils"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/log"
	"github.com/FusionFoundation/fsn-go-sdk/mongosync/syncer"
	"gopkg.in/urfave/cli.v1"
)

var (
	mongoURLFlag = cli.StringFlag{
		Name:  "mongoURL",
		Usage: "mongodb URL",
		Value: "localhost:27017",
	}
	dbNameFlag = cli.StringFlag{
		Name:  "dbName",
		Usage: "database name",
		Value: "fusion",
	}
	stableFlag = cli.Uint64Flag{
		Name:  "stable",
		Usage: "sync blocks that is stable lower than the latest",
		Value: 10,
	}
	startFlag = cli.Uint64Flag{
		Name:  "start",
		Usage: "sync blocks from this height",
		Value: 0,
	}
	endFlag = cli.Uint64Flag{
		Name:  "end",
		Usage: "sync blocks to this height (not inclusive)",
		Value: 0,
	}
	overwriteFlag = cli.BoolFlag{
		Name:  "overwrite",
		Usage: "overwrite synced entries",
	}
	archiveModeFlag = cli.BoolTFlag{
		Name:  "archive",
		Usage: "specify whether the server node is using archive mode. If not we may not get retreated tickets info etc.",
	}
)

var commandSyncAll = cli.Command{
	Name:      "syncAll",
	Usage:     "sync all info into mongodb",
	ArgsUsage: "<severURL>",
	Description: `
Sync all info into mongodb.
severURL support "http", "https", "ws", "wss", "stdio", IPC file`,
	Flags: []cli.Flag{
		mongoURLFlag,
		dbNameFlag,
		stableFlag,
		startFlag,
		endFlag,
		overwriteFlag,
		archiveModeFlag,
	},
	Action: syncAll,
}

func syncAll(ctx *cli.Context) error {
	log.SetLogger(ctx.GlobalInt(utils.VerbosityFlag.Name), ctx.GlobalBool(utils.JsonFlag.Name))
	if len(ctx.Args()) == 0 {
		cli.ShowCommandHelpAndExit(ctx, "syncAll", 1)
	}

	serverAddress := ctx.Args().First()
	mongoURL := ctx.String(mongoURLFlag.Name)
	dbName := ctx.String(dbNameFlag.Name)
	stable := ctx.Uint64(stableFlag.Name)
	start := ctx.Uint64(startFlag.Name)
	end := ctx.Uint64(endFlag.Name)
	overwrite := ctx.Bool(overwriteFlag.Name)
	archiveMode := ctx.Bool(archiveModeFlag.Name)
	jobs := ctx.GlobalUint64(jobsFlag.Name)
	interval := ctx.GlobalUint64(intervalFlag.Name)

	syncer.ServerURL = serverAddress
	syncer.Overwrite = overwrite
	syncer.ArchiveMode = archiveMode
	syncer.BlockInterval = interval
	syncer.SetJobCount(jobs)
	syncer.InitMongoServer(mongoURL, dbName)
	syncer.NewSyncer(stable, start, end).Sync()

	return nil
}
