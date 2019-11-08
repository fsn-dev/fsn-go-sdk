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
	"fmt"
	"math/big"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/cmd/utils"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/common"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/common/hexutil"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/core/types"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/rlp"
	clicommon "github.com/FusionFoundation/fsn-go-sdk/fsn-cli/common"
	"github.com/FusionFoundation/fsn-go-sdk/fsnapi"
	"gopkg.in/urfave/cli.v1"
)

var (
	signFlag = cli.BoolFlag{
		Name:  "sign",
		Usage: "sign the transaction",
	}
	senderFlag = cli.StringFlag{
		Name:  "from",
		Usage: "transaction sender",
		Value: "",
	}
	accountNonceFlag = cli.Uint64Flag{
		Name:  "nonce",
		Usage: "set account nonce",
		Value: 0,
	}
	gasLimitFlag = cli.Uint64Flag{
		Name:  "gaslimit",
		Usage: "set gas limit",
		Value: 90000,
	}
	gasPriceFlag = cli.StringFlag{
		Name:  "gasprice",
		Usage: "set gas price",
		Value: "1000000000",
	}
	keyStoreFileFlag = cli.StringFlag{
		Name:  "keystore",
		Usage: "keystore file to use for signing transaction",
		Value: "",
	}

	commonFlags = []cli.Flag{
		signFlag,
		senderFlag,
		accountNonceFlag,
		gasLimitFlag,
		gasPriceFlag,
		keyStoreFileFlag,
		utils.PasswordFileFlag,
	}
)

func getHexUint64(ctx *cli.Context, flagName string) *hexutil.Uint64 {
	value := ctx.Uint64(flagName)
	result := new(hexutil.Uint64)
	*(*uint64)(result) = value
	return result
}

func getHexBigInt(ctx *cli.Context, flagName string) *hexutil.Big {
	value := clicommon.GetBigIntFromText(flagName, ctx.String(flagName))
	result := new(hexutil.Big)
	*(*big.Int)(result) = *value
	return result
}

func getHash(ctx *cli.Context, flagName string) common.Hash {
	return clicommon.GetHashFromText(flagName, ctx.String(flagName))
}

func getAddress(ctx *cli.Context, flagName string) common.Address {
	return clicommon.GetAddressFromText(flagName, ctx.String(flagName))
}

func getBigInt(ctx *cli.Context, flagName string) *big.Int {
	return clicommon.GetBigIntFromText(flagName, ctx.String(flagName))
}

func getBaseArgsAndSignOptions(ctx *cli.Context) (common.FusionBaseArgs, *fsnapi.SignOptions) {
	var (
		args     common.FusionBaseArgs
		signopts *fsnapi.SignOptions

		from common.Address
	)

	if ctx.IsSet(senderFlag.Name) {
		from = clicommon.GetAddressFromText("from", ctx.String(senderFlag.Name))
		args.From = from
	}

	if ctx.IsSet(accountNonceFlag.Name) {
		args.Nonce = getHexUint64(ctx, accountNonceFlag.Name)
	}

	args.Gas = getHexUint64(ctx, gasLimitFlag.Name)
	args.GasPrice = getHexBigInt(ctx, gasPriceFlag.Name)

	if ctx.Bool(signFlag.Name) {
		signopts = &fsnapi.SignOptions{}
		signopts.Signer = from
		signopts.Keyfile = ctx.String(keyStoreFileFlag.Name)
		signopts.Passfile = ctx.String(utils.PasswordFileFlag.Name)
		if args.From == (common.Address{}) ||
			args.Nonce == nil ||
			signopts.Keyfile == "" ||
			signopts.Passfile == "" {
			utils.Fatalf("Must provide (--%s --%s --%s --%s) options to sign transaction", senderFlag.Name, accountNonceFlag.Name, keyStoreFileFlag.Name, utils.PasswordFileFlag.Name)
		}
	}

	return args, signopts
}

func printTx(tx *types.Transaction, json bool) error {
	if json {
		bs, err := tx.MarshalJSONWithSender()
		if err != nil {
			return fmt.Errorf("json marshal err %v", err)
		}
		fmt.Println(string(bs))
	} else {
		bs, err := rlp.EncodeToBytes(tx)
		if err != nil {
			return fmt.Errorf("rlp encode err %v", err)
		}
		fmt.Println(hexutil.Bytes(bs))
	}
	return nil
}
