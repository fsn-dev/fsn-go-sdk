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

	"github.com/fsn-dev/fsn-go-sdk/efsn/tools"
	clicommon "github.com/fsn-dev/fsn-go-sdk/fsn-cli/common"
	"gopkg.in/urfave/cli.v1"
)

var CommandGetTransactionCount = cli.Command{
	Name:      "gettransactioncount",
	Aliases:   []string{"getnonce"},
	Category:  "online",
	Usage:     "get transaction count",
	ArgsUsage: "<address>",
	Description: `
get transaction count (account nonce)`,
	Flags: []cli.Flag{
		serverAddrFlag,
	},
	Action: gettransactioncount,
}

func gettransactioncount(ctx *cli.Context) error {
	setLogger(ctx)
	if len(ctx.Args()) != 1 {
		cli.ShowCommandHelpAndExit(ctx, "gettransactioncount", 1)
	}

	client := dialServer(ctx)
	defer client.Close()

	address := clicommon.GetAddressFromText("txHash", ctx.Args().First())
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return err
	}

	tools.MustPrintJSON(nonce)
	return nil
}
