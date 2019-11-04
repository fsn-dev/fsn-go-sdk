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
	"bytes"
	"context"
	"fmt"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/common/hexutil"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/core/types"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/rlp"
	"gopkg.in/urfave/cli.v1"
)

var CommandSendRawTx = cli.Command{
	Name:      "sendrawtx",
	Usage:     "(online) send raw transaction",
	ArgsUsage: "<rawtx>",
	Description: `
broadcast offline signed raw transaction (hex encoded)`,
	Flags: []cli.Flag{
		serverAddrFlag,
	},
	Action: sendrawtx,
}

func sendrawtx(ctx *cli.Context) (err error) {
	if len(ctx.Args()) != 1 {
		cli.ShowCommandHelpAndExit(ctx, "sendrawtx", 1)
	}

	client := dialServer(ctx)
	defer client.Close()

	rawtx := ctx.Args().First()
	data, err := hexutil.Decode(rawtx)
	if err != nil {
		return fmt.Errorf("wrong arguments %v", err)
	}

	var tx types.Transaction
	err = rlp.Decode(bytes.NewReader(data), &tx)
	if err != nil {
		return fmt.Errorf("decode raw transaction err %v", err)
	}

	err = client.SendTransaction(context.Background(), &tx)
	if err != nil {
		return err
	}

	fmt.Println(tx.Hash())
	return nil
}
