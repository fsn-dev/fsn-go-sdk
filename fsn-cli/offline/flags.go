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
	ticketStartFlag = cli.Uint64Flag{
		Name:  "start",
		Usage: "ticket start time, 0 means now",
		Value: 0,
	}
	ticketEndFlag = cli.Uint64Flag{
		Name:  "end",
		Usage: "ticket end time, 0 means start + 1 month",
		Value: 0,
	}

	nameFlag = cli.StringFlag{
		Name:  "name",
		Usage: "",
	}
	symbolFlag = cli.StringFlag{
		Name:  "symbol",
		Usage: "",
	}
	totalFlag = cli.StringFlag{
		Name:  "total",
		Usage: "",
	}
	decimalsFlag = cli.Uint64Flag{
		Name:  "decimals",
		Usage: "",
		Value: 0,
	}
	canChangeFlag = cli.BoolFlag{
		Name:  "canchange",
		Usage: "can change total supply",
	}
	descriptionFlag = cli.StringFlag{
		Name:  "description",
		Usage: "",
	}
)
