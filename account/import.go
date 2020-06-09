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
	"fmt"

	"github.com/fsn-dev/fsn-go-sdk/efsn/accounts/keystore"
	"github.com/fsn-dev/fsn-go-sdk/efsn/cmd/utils"
	"github.com/fsn-dev/fsn-go-sdk/efsn/crypto"
	"gopkg.in/urfave/cli.v1"
)

var commandImport = cli.Command{
	Name:   "import",
	Usage:  "Import a private key into a new account",
	Action: accountImport,
	Flags: []cli.Flag{
		utils.DataDirFlag,
		utils.KeyStoreDirFlag,
		utils.PasswordFileFlag,
	},
	ArgsUsage: "<keyFile>",
	Description: `
    efsn account import <keyfile>

Imports an unencrypted private key from <keyfile> and creates a new account.
Prints the address.

The keyfile is assumed to contain an unencrypted private key in hexadecimal format.

The account is saved in encrypted format, you are prompted for a passphrase.

You must remember this passphrase to unlock your account in the future.

For non-interactive use the passphrase can be specified with the -password flag:

    efsn account import [options] <keyfile>

Note:
As you can directly copy your encrypted accounts to another ethereum instance,
this import mechanism is not needed when you transfer an account between
nodes.
`,
}

func accountImport(ctx *cli.Context) error {
	keyfile := ctx.Args().First()
	if len(keyfile) == 0 {
		utils.Fatalf("keyfile must be given as argument")
	}
	key, err := crypto.LoadECDSA(keyfile)
	if err != nil {
		utils.Fatalf("Failed to load the private key: %v", err)
	}

	cfg := &utils.Config{}
	utils.SetNodeConfig(ctx, cfg)
	am, _, err := cfg.AccountManager()
	if err != nil {
		utils.Fatalf("Failed to get account manager: %v", err)
	}

	passphrase := getPassPhrase("Your new account is locked with a password. Please give a password. Do not forget this password.", true, 0, utils.MakePasswordList(ctx))

	ks := am.Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	acct, err := ks.ImportECDSA(key, passphrase)
	if err != nil {
		utils.Fatalf("Could not create the account: %v", err)
	}
	fmt.Printf("Address: {%x}\n", acct.Address)
	return nil
}
