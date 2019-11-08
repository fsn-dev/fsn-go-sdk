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
