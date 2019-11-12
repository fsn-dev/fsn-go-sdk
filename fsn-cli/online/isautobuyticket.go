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
	"gopkg.in/urfave/cli.v1"
)

var CommandIsAutoBuyTicket = cli.Command{
	Name:      "isautobuyticket",
	Aliases:   []string{"isautobt"},
	Usage:     "(online) get is auto buy ticket flag",
	ArgsUsage: "",
	Description: `
get is auto buy ticket flag`,
	Flags: []cli.Flag{
		serverAddrFlag,
	},
	Action: isautobuyticket,
}

func isautobuyticket(ctx *cli.Context) error {
	setLogger(ctx)

	client := dialServer(ctx)
	defer client.Close()

	flag, err := client.IsAutoBuyTicket(context.Background())
	if err != nil {
		return err
	}

	tools.MustPrintJSON(flag)
	return nil
}
