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
	"math/big"
	"time"

	"github.com/fsn-dev/fsn-go-sdk/efsn/cmd/utils"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common/hexutil"
	"github.com/fsn-dev/fsn-go-sdk/efsn/core/types"
	"github.com/fsn-dev/fsn-go-sdk/efsn/log"
	"github.com/fsn-dev/fsn-go-sdk/efsn/tools"
	clicommon "github.com/fsn-dev/fsn-go-sdk/fsn-cli/common"
	"github.com/fsn-dev/fsn-go-sdk/fsnapi"
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
	}
	accountNonceFlag = cli.Uint64Flag{
		Name:  "nonce",
		Usage: "set account nonce",
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
		Name:   "keystore",
		Usage:  "keystore file to use for signing transaction",
		EnvVar: "FSN_KEYSTORE_FILE",
	}
	chainIdFlag = cli.Uint64Flag{
		Name:  "chainid",
		Usage: "chain identifier (integer, 32659=mainnet, 46688=testnet, 55555=devnet)",
		Value: 32659,
	}

	signTxFlags = []cli.Flag{
		senderFlag,
		accountNonceFlag,
		gasLimitFlag,
		gasPriceFlag,
		keyStoreFileFlag,
		utils.PasswordFileFlag,
		chainIdFlag,
	}

	commonFlags = append([]cli.Flag{signFlag}, signTxFlags...)
)

func setLogger(ctx *cli.Context) {
	log.SetLogger(ctx.GlobalInt(utils.VerbosityFlag.Name), ctx.GlobalBool(utils.JsonFlag.Name))
}

func getNowTime() uint64 {
	return uint64(time.Now().Unix())
}

func getBigIntFromUint64(num uint64) *big.Int {
	return new(big.Int).SetUint64(num)
}

func getHexUint64(ctx *cli.Context, flagName string) *hexutil.Uint64 {
	value := ctx.Uint64(flagName)
	result := new(hexutil.Uint64)
	*(*uint64)(result) = value
	return result
}

func getHexUint64Time(ctx *cli.Context, flagName string) *hexutil.Uint64 {
	if !ctx.IsSet(flagName) {
		return nil
	}
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

func getAddressSlice(ctx *cli.Context, flagName string) []common.Address {
	return clicommon.GetAddressSlice(flagName, ctx.StringSlice(flagName))
}

func getHashSlice(ctx *cli.Context, flagName string) []common.Hash {
	return clicommon.GetHashSlice(flagName, ctx.StringSlice(flagName))
}

func getHexBigIntSlice(ctx *cli.Context, flagName string) []*hexutil.Big {
	return clicommon.GetHexBigIntSlice(flagName, ctx.StringSlice(flagName))
}

func getHexUint64Slice(ctx *cli.Context, flagName string) []*hexutil.Uint64 {
	return clicommon.GetHexUint64Slice(flagName, ctx.Int64Slice(flagName))
}

func getBaseArgsAndSignOptions(ctx *cli.Context) (common.FusionBaseArgs, *fsnapi.SignOptions) {
	return getBaseArgsAndSignOptionsImpl(ctx, false)
}

func getBaseArgsAndSignOptionsForSign(ctx *cli.Context) (common.FusionBaseArgs, *fsnapi.SignOptions) {
	return getBaseArgsAndSignOptionsImpl(ctx, true)
}

func getBaseArgsAndSignOptionsImpl(ctx *cli.Context, forceSign bool) (common.FusionBaseArgs, *fsnapi.SignOptions) {
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

	if forceSign || ctx.Bool(signFlag.Name) {
		signopts = &fsnapi.SignOptions{}
		signopts.Signer = from
		signopts.Keyfile = ctx.String(keyStoreFileFlag.Name)
		signopts.ChainID = ctx.Uint64(chainIdFlag.Name)
		if signopts.Keyfile == "" {
			utils.Fatalf("must specify '%s' option or set '%s' enviroment", keyStoreFileFlag.Name, keyStoreFileFlag.EnvVar)
		}
		if args.From == (common.Address{}) || args.Nonce == nil {
			utils.Fatalf("Must provide ('%s', '%s') options to sign transaction", senderFlag.Name, accountNonceFlag.Name)
		}
		if ctx.IsSet(utils.PasswordFileFlag.Name) {
			signopts.Passfile = ctx.String(utils.PasswordFileFlag.Name)
		} else {
			password, err := tools.Stdin.PromptPassword("Passphrase: ")
			if err != nil {
				utils.Fatalf("Failed to read passphrase: %v", err)
			}
			signopts.Password = password
		}
	}

	return args, signopts
}

func printTx(tx *types.Transaction, json bool) error {
	return clicommon.PrintTx(tx, json)
}
