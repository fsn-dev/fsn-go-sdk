// extend ethclient for fusion
package ethclient

import (
	"context"
	"math/big"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/common"
)

func (ec *Client) GetAsset(ctx context.Context, assetId common.Hash, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getAsset", assetId.String(), toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetTimeLockBalance(ctx context.Context, assetId common.Hash, address common.Address, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getTimeLockBalance", assetId.String(), address, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetBalance(ctx context.Context, assetId common.Hash, address common.Address, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getBalance", assetId.String(), address, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}
