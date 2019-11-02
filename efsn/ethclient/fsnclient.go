// extend ethclient for fusion
package ethclient

import (
	"context"
	"math/big"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/common"
)

type FSNAsset struct {
	ID          common.Hash
	Owner       common.Address
	Name        string
	Symbol      string
	Decimals    uint8
	Total       string
	CanChange   bool
	Description string
}

func (ec *Client) GetAsset(ctx context.Context, assetId common.Hash, blockNr *big.Int) (*FSNAsset, error) {
	var asset FSNAsset
	err := ec.c.CallContext(ctx, &asset, "fsn_getAsset", assetId.String(), toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &asset, nil
}
