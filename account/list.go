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

	"github.com/fsn-dev/fsn-go-sdk/efsn/cmd/utils"
	"gopkg.in/urfave/cli.v1"
)

var commandList = cli.Command{
	Name:   "list",
	Usage:  "Print summary of existing accounts",
	Action: accountList,
	Flags: []cli.Flag{
		utils.DataDirFlag,
		utils.KeyStoreDirFlag,
	},
	Description: `
Print a short summary of all accounts`,
}

func accountList(ctx *cli.Context) error {
	cfg := &utils.Config{}
	utils.SetNodeConfig(ctx, cfg)
	am, _, err := cfg.AccountManager()
	if err != nil {
		utils.Fatalf("Failed to get account manager: %v", err)
	}
	var index int
	for _, wallet := range am.Wallets() {
		for _, account := range wallet.Accounts() {
			fmt.Printf("Account #%d: {%x} %s\n", index, account.Address, &account.URL)
			index++
		}
	}
	return nil
}
