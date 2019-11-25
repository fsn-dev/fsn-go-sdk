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

package online

import (
	"context"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/tools"
	clicommon "github.com/FusionFoundation/fsn-go-sdk/fsn-cli/common"
	"gopkg.in/urfave/cli.v1"
)

var CommandGetStakeInfo = cli.Command{
	Name:      "getstakeinfo",
	Aliases:   []string{"getstake"},
	Category:  "online",
	Usage:     "get stake info",
	ArgsUsage: "",
	Description: `
get stake information`,
	Flags: []cli.Flag{
		blockHeightFlag,
		serverAddrFlag,
	},
	Action: getstakeinfo,
}

func getstakeinfo(ctx *cli.Context) error {
	setLogger(ctx)

	client := dialServer(ctx)
	defer client.Close()

	blockNr := clicommon.GetBlockNumberFromText(ctx.String(blockHeightFlag.Name))
	info, err := client.GetStakeInfo(context.Background(), blockNr)
	if err != nil {
		return err
	}

	tools.MustPrintJSON(info)
	return nil
}
