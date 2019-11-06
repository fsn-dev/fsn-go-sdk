package query

import (
	"context"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/common"
	"math/big"
)

func (c client) GetAccount(addr string)(*big.Int,error){
	return c.getAccount(addr,nil)
}

func (c client) GetAccountAtBlockNumber(addr string,blockNumber int64)(*big.Int,error){
	blockNumberInt := big.NewInt(blockNumber)
	return c.getAccount(addr,blockNumberInt)
}

func (c client) getAccount(addrStr string,blockNumber *big.Int)(*big.Int,error){
	addr := common.HexToAddress(addrStr)
	ctx := context.Background()
	return c.BalanceAt(ctx,addr,blockNumber)
}
