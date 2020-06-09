// Copyright 2016 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/fsn-dev/fsn-go-sdk/efsn/accounts/keystore"
	"github.com/fsn-dev/fsn-go-sdk/efsn/cmd/utils"
	"github.com/fsn-dev/fsn-go-sdk/efsn/crypto"
	"gopkg.in/urfave/cli.v1"
)

var commandShowkey = cli.Command{
	Name:   "showkey",
	Usage:  "Show private key of a keystore file",
	Action: accountShowkey,
	Flags: []cli.Flag{
		utils.PasswordFileFlag,
	},
	ArgsUsage: "<keyFile>",
	Description: `
Show private key of a keystore file

For non-interactive use the passphrase can be specified with the --password flag:

    efsn account showkey [options] <keyFile>
`,
}

func accountShowkey(ctx *cli.Context) error {
	keyfile := ctx.Args().First()
	if len(keyfile) == 0 {
		utils.Fatalf("keyfile must be given as argument")
	}
	keyJSON, err := ioutil.ReadFile(keyfile)
	if err != nil {
		utils.Fatalf("Could not read wallet file: %v", err)
	}

	// Decrypt key with passphrase.
	passphrase := getPassPhrase("", false, 0, utils.MakePasswordList(ctx))
	key, err := keystore.DecryptKey(keyJSON, passphrase)
	if err != nil {
		utils.Fatalf("Error decrypting key: %v", err)
	}
	privateKey := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))

	fmt.Printf("Address: {%s}, PrivateKey: {%s}\n", key.Address.String(), privateKey)
	return nil
}
