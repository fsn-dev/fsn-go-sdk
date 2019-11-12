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

func (ec *Client) GetBalance(ctx context.Context, assetId common.Hash, address common.Address, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getBalance", assetId.String(), address, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetAllBalances(ctx context.Context, address common.Address, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getAllBalances", address, toBlockNumArg(blockNr))
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

func (ec *Client) GetRawTimeLockBalance(ctx context.Context, assetId common.Hash, address common.Address, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getRawTimeLockBalance", assetId.String(), address, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetAllTimeLockBalances(ctx context.Context, address common.Address, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getAllTimeLockBalances", address, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetAllRawTimeLockBalances(ctx context.Context, address common.Address, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getAllRawTimeLockBalances", address, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetSnapshot(ctx context.Context, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getSnapshot", toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetSnapshotAtHash(ctx context.Context, hash common.Hash) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getSnapshotAtHash", hash.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetSwap(ctx context.Context, swapId common.Hash, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getSwap", swapId.String(), toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetMultiSwap(ctx context.Context, swapId common.Hash, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getMultiSwap", swapId.String(), toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetNotation(ctx context.Context, address common.Address, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getNotation", address.String(), toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetLatestNotation(ctx context.Context, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getLatestNotation", toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetAddressByNotation(ctx context.Context, notation uint64, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getAddressByNotation", notation, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) AllTickets(ctx context.Context, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_allTickets", toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) AllTicketsByAddress(ctx context.Context, address common.Address, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_allTicketsByAddress", address.String(), toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) IsAutoBuyTicket(ctx context.Context) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_isAutoBuyTicket")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) TicketPrice(ctx context.Context, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_ticketPrice", toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) TotalNumberOfTickets(ctx context.Context, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_totalNumberOfTickets", toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) TotalNumberOfTicketsByAddress(ctx context.Context, address common.Address, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_totalNumberOfTicketsByAddress", address.String(), toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) AllInfoByAddress(ctx context.Context, address common.Address, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_allInfoByAddress", address.String(), toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetTransactionAndReceipt(ctx context.Context, hash common.Hash) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getTransactionAndReceipt", hash.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetStakeInfo(ctx context.Context, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getStakeInfo", toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetBlockReward(ctx context.Context, blockNr *big.Int) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "fsn_getBlockReward", toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetTransaction(ctx context.Context, txHash common.Hash) (*interface{}, error) {
	var result interface{}
	err := ec.c.CallContext(ctx, &result, "eth_getTransactionByHash", txHash)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
