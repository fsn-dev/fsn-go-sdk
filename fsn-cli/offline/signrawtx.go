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
	"github.com/fsn-dev/fsn-go-sdk/efsn/core/types"
	"github.com/fsn-dev/fsn-go-sdk/fsnapi"
	"gopkg.in/urfave/cli.v1"
)

var CommandSignRawTx = cli.Command{
	Name:      "signrawtx",
	Aliases:   []string{"signtx"},
	Category:  "offline",
	Usage:     "sign raw transaction",
	ArgsUsage: "<hexdata>",
	Description: `
sign raw transaction, you can replace corresponding arguments (eg. gas price)`,
	Flags:  signTxFlags,
	Action: signrawtx,
}

func signrawtx(ctx *cli.Context) error {
	setLogger(ctx)
	if len(ctx.Args()) != 1 {
		cli.ShowCommandHelpAndExit(ctx, "signrawtx", 1)
	}

	hexstr := ctx.Args().First()
	tx, err := fsnapi.DecodeRawTx(hexstr)
	if err != nil {
		return err
	}

	baseArgs, signOptions := getBaseArgsAndSignOptionsForSign(ctx)

	nonce := tx.Nonce()
	gasLimit := tx.Gas()
	gasPrice := tx.GasPrice()

	if baseArgs.Nonce != nil {
		nonce = uint64(*baseArgs.Nonce)
	}
	if baseArgs.Gas != nil {
		gasLimit = uint64(*baseArgs.Gas)
	}
	if baseArgs.GasPrice != nil {
		gasPrice = baseArgs.GasPrice.ToInt()
	}

	if tx.To() != nil {
		tx = types.NewTransaction(
			nonce,
			*tx.To(),
			tx.Value(),
			gasLimit,
			gasPrice,
			tx.Data(),
		)
	} else {
		tx = types.NewContractCreation(
			nonce,
			tx.Value(),
			gasLimit,
			gasPrice,
			tx.Data(),
		)
	}

	tx, err = fsnapi.SignTx(tx, signOptions)
	if err != nil {
		return err
	}
	return printTx(tx, false)
}
