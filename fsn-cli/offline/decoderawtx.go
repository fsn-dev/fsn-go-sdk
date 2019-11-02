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
	"bytes"
	"fmt"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/common/hexutil"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/core/types"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/rlp"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/tools"
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
	data, err := hexutil.Decode(hexstr)
	if err != nil {
		return fmt.Errorf("wrong arguments %v", err)
	}

	if ctx.Bool(decodeTxInputFlag.Name) {
		return decodeTxInput(data)
	}

	if ctx.Bool(decodeLogDataFlag.Name) {
		return decodeLogData(data)
	}

	var tx types.Transaction
	err = rlp.Decode(bytes.NewReader(data), &tx)
	if err != nil {
		return fmt.Errorf("decode raw transaction err %v", err)
	}
	return printTx(&tx, true)
}

func decodeTxInput(data []byte) error {
	res, err := tools.DecodeTxInput(data)
	if err != nil {
		return fmt.Errorf("decode transaction input err %v", err)
	}
	tools.MustPrintJSON(res)
	return nil
}

func decodeLogData(data []byte) error {
	res, err := tools.DecodeLogData(data)
	if err != nil {
		return fmt.Errorf("decode log data err %v", err)
	}
	tools.MustPrintJSON(res)
	return nil
}
