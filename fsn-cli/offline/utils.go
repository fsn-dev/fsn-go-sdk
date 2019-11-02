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
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/accounts/keystore"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/cmd/utils"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/common"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/common/hexutil"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/core/types"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/params"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/rlp"
	clicommon "github.com/FusionFoundation/fsn-go-sdk/fsn-cli/common"
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

type CommonTxOptions struct {
	nonce    uint64
	gasLimit uint64
	gasPrice *big.Int
}

type CommonSignOptions struct {
	signer   common.Address
	keyfile  string
	passfile string
}

type CommonOptions struct {
	needSign bool
	*CommonTxOptions
	*CommonSignOptions
}

func getCommonOptions(ctx *cli.Context) *CommonOptions {
	var (
		needSign bool
		txopts   CommonTxOptions
		signopts CommonSignOptions
	)

	needSign = ctx.Bool(signFlag.Name)

	txopts.nonce = ctx.Uint64(accountNonceFlag.Name)
	txopts.gasLimit = ctx.Uint64(gasLimitFlag.Name)
	txopts.gasPrice = clicommon.GetBigIntFromText("gasPrice", ctx.String(gasPriceFlag.Name))

	if needSign {
		signopts.signer = clicommon.GetAddressFromText("from", ctx.String(senderFlag.Name))
		signopts.keyfile = ctx.String(keyStoreFileFlag.Name)
		signopts.passfile = ctx.String(utils.PasswordFileFlag.Name)
	}

	return &CommonOptions{
		needSign:          needSign,
		CommonTxOptions:   &txopts,
		CommonSignOptions: &signopts,
	}
}

func printTx(tx *types.Transaction, json bool) error {
	if json {
		bs, err := tx.MarshalJSON()
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

type ParamInterface interface {
	ToBytes() ([]byte, error)
}

func genTxInput(funcType common.FSNCallFunc, funcParam ParamInterface) ([]byte, error) {
	funcData, err := funcParam.ToBytes()
	if err != nil {
		return nil, err
	}
	var param = common.FSNCallParam{Func: funcType, Data: funcData}
	input, err := param.ToBytes()
	if err != nil {
		return nil, err
	}
	return input, nil
}

func toFSNTx(input []byte, commonOptions *CommonOptions) (tx *types.Transaction, err error) {
	tx = types.NewTransaction(
		commonOptions.nonce,
		common.FSNCallAddress,
		big.NewInt(0),
		commonOptions.gasLimit,
		commonOptions.gasPrice,
		input,
	)
	if commonOptions.needSign {
		tx, err = signTx(tx, commonOptions.CommonSignOptions)
		if err != nil {
			return nil, err
		}
	}
	return tx, nil
}

func signTx(tx *types.Transaction, signOptions *CommonSignOptions) (*types.Transaction, error) {
	keyjson, err := ioutil.ReadFile(signOptions.keyfile)
	if err != nil {
		return nil, err
	}

	passdata, err := ioutil.ReadFile(signOptions.passfile)
	if err != nil {
		return nil, err
	}

	passphrase := strings.TrimSpace(string(passdata))
	key, err := keystore.DecryptKey(keyjson, passphrase)
	if err != nil {
		return nil, err
	}

	if key.Address != signOptions.signer {
		return nil, fmt.Errorf("key content mismatch: have account %x, want %x", key.Address, signOptions.signer)
	}

	signer := types.NewEIP155Signer(params.MainnetChainConfig.ChainID)
	return types.SignTx(tx, signer, key.PrivateKey)
}
