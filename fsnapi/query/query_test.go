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

package query

import (
	"fmt"
	"testing"
)

func TestClient_GetAccount(t *testing.T) {
	c, err := NewClient("http://127.0.0:7771")
	if err != nil {
		panic(err)
	}
	balance1, err := c.GetAccount("0x9383FcC878e587a84d25E1ab956145360c0F82F3")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(balance1)

	balance2, err := c.GetAccountAtBlockNumber("0x9383FcC878e587a84d25E1ab956145360c0F82F3", 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(balance2)

}
