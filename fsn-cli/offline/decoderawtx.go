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
	"github.com/FusionFoundation/fsn-go-sdk/efsn/tools"
	"github.com/FusionFoundation/fsn-go-sdk/fsnapi"
	"gopkg.in/urfave/cli.v1"
)

var CommandDecodeRawTx = cli.Command{
	Name:      "decoderawtx",
	Usage:     "(offline) decode raw transaction data",
	ArgsUsage: "<hexdata>",
	Description: `
decode raw transaction data`,
	Flags: []cli.Flag{
		decodeTxInputFlag,
		decodeLogDataFlag,
	},
	Action: decoderawtx,
}

func decoderawtx(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		cli.ShowCommandHelpAndExit(ctx, "decoderawtx", 1)
	}

	hexstr := ctx.Args().First()

	if ctx.Bool(decodeTxInputFlag.Name) {
		return decodeTxInput(hexstr)
	}

	if ctx.Bool(decodeLogDataFlag.Name) {
		return decodeTxReceiptLogData(hexstr)
	}

	tx, err := fsnapi.DecodeRawTx(hexstr)
	if err != nil {
		return err
	}
	return printTx(tx, true)
}

func decodeTxInput(hexstr string) error {
	res, err := fsnapi.DecodeTxInput(hexstr)
	if err != nil {
		return err
	}
	tools.MustPrintJSON(res)
	return nil
}

func decodeTxReceiptLogData(hexstr string) error {
	res, err := fsnapi.DecodeTxReceiptLogData(hexstr)
	if err != nil {
		return err
	}
	tools.MustPrintJSON(res)
	return nil
}
