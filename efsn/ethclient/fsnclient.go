// extend ethclient for fusion
package ethclient

import (
	"context"
	"math/big"

	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common/hexutil"
	"github.com/fsn-dev/fsn-go-sdk/efsn/tools"
)

func strToBigInt(str string) *big.Int {
	bi, ok := new(big.Int).SetString(str, 0)
	if ok {
		return bi
	}
	return nil
}

type RPCAsset struct {
	ID          common.Hash
	Owner       common.Address
	Name        string
	Symbol      string
	Decimals    uint8
	Total       string
	CanChange   bool
	Description string
}

func (r *RPCAsset) ToAsset() *common.Asset {
	return &common.Asset{
		ID:          r.ID,
		Owner:       r.Owner,
		Name:        r.Name,
		Symbol:      r.Symbol,
		Decimals:    r.Decimals,
		Total:       strToBigInt(r.Total),
		CanChange:   r.CanChange,
		Description: r.Description,
	}
}

func (ec *Client) GetAsset(ctx context.Context, assetId common.Hash, blockNr *big.Int) (*common.Asset, error) {
	var result RPCAsset
	err := ec.c.CallContext(ctx, &result, "fsn_getAsset", assetId.String(), toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return result.ToAsset(), nil
}

func (ec *Client) GetBalance(ctx context.Context, assetId common.Hash, address common.Address, blockNr *big.Int) (*big.Int, error) {
	var result string
	err := ec.c.CallContext(ctx, &result, "fsn_getBalance", assetId.String(), address, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return strToBigInt(result), nil
}

func (ec *Client) GetAllBalances(ctx context.Context, address common.Address, blockNr *big.Int) (map[common.Hash]*big.Int, error) {
	var result map[common.Hash]string
	err := ec.c.CallContext(ctx, &result, "fsn_getAllBalances", address, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	ret := make(map[common.Hash]*big.Int, len(result))
	for k, v := range result {
		ret[k] = strToBigInt(v)
	}
	return ret, nil
}

type RPCTimeLockItem struct {
	StartTime uint64
	EndTime   uint64
	Value     string
}
type RPCTimeLock struct {
	Items []*RPCTimeLockItem
}

func (r *RPCTimeLockItem) ToTimeLockItem() *common.TimeLockItem {
	return &common.TimeLockItem{
		StartTime: r.StartTime,
		EndTime:   r.EndTime,
		Value:     strToBigInt(r.Value),
	}
}

func (r *RPCTimeLock) ToTimeLock() *common.TimeLock {
	ret := &common.TimeLock{
		Items: make([]*common.TimeLockItem, len(r.Items)),
	}
	for i, item := range r.Items {
		ret.Items[i] = item.ToTimeLockItem()
	}
	return ret
}

func (ec *Client) GetTimeLockBalance(ctx context.Context, assetId common.Hash, address common.Address, blockNr *big.Int) (*common.TimeLock, error) {
	var result RPCTimeLock
	err := ec.c.CallContext(ctx, &result, "fsn_getTimeLockBalance", assetId.String(), address, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return result.ToTimeLock(), nil
}

func (ec *Client) GetTimeLockValueByInterval(ctx context.Context, assetId common.Hash, address common.Address, startTime, endTime uint64, blockNr *big.Int) (*big.Int, error) {
	var result string
	err := ec.c.CallContext(ctx, &result, "fsn_getTimeLockValueByInterval", assetId.String(), address, startTime, endTime, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return strToBigInt(result), nil
}

func (ec *Client) GetRawTimeLockBalance(ctx context.Context, assetId common.Hash, address common.Address, blockNr *big.Int) (*common.TimeLock, error) {
	var result RPCTimeLock
	err := ec.c.CallContext(ctx, &result, "fsn_getRawTimeLockBalance", assetId.String(), address, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return result.ToTimeLock(), nil
}

func (ec *Client) GetAllTimeLockBalances(ctx context.Context, address common.Address, blockNr *big.Int) (map[common.Hash]*common.TimeLock, error) {
	var result map[common.Hash]*RPCTimeLock
	err := ec.c.CallContext(ctx, &result, "fsn_getAllTimeLockBalances", address, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	ret := make(map[common.Hash]*common.TimeLock, len(result))
	for k, v := range result {
		ret[k] = v.ToTimeLock()
	}
	return ret, nil
}

func (ec *Client) GetAllRawTimeLockBalances(ctx context.Context, address common.Address, blockNr *big.Int) (map[common.Hash]*common.TimeLock, error) {
	var result map[common.Hash]*RPCTimeLock
	err := ec.c.CallContext(ctx, &result, "fsn_getAllRawTimeLockBalances", address, toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	ret := make(map[common.Hash]*common.TimeLock, len(result))
	for k, v := range result {
		ret[k] = v.ToTimeLock()
	}
	return ret, nil
}

func (ec *Client) GetSnapshot(ctx context.Context, blockNr *big.Int) (*tools.Snapshot, error) {
	var result tools.Snapshot
	err := ec.c.CallContext(ctx, &result, "fsn_getSnapshot", toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetSnapshotAtHash(ctx context.Context, hash common.Hash) (*tools.Snapshot, error) {
	var result tools.Snapshot
	err := ec.c.CallContext(ctx, &result, "fsn_getSnapshotAtHash", hash.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetSwap(ctx context.Context, swapId common.Hash, blockNr *big.Int) (*common.Swap, error) {
	var result common.Swap
	err := ec.c.CallContext(ctx, &result, "fsn_getSwap", swapId.String(), toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetMultiSwap(ctx context.Context, swapId common.Hash, blockNr *big.Int) (*common.MultiSwap, error) {
	var result common.MultiSwap
	err := ec.c.CallContext(ctx, &result, "fsn_getMultiSwap", swapId.String(), toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetNotation(ctx context.Context, address common.Address, blockNr *big.Int) (uint64, error) {
	var result uint64
	err := ec.c.CallContext(ctx, &result, "fsn_getNotation", address.String(), toBlockNumArg(blockNr))
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (ec *Client) GetLatestNotation(ctx context.Context, blockNr *big.Int) (uint64, error) {
	var result uint64
	err := ec.c.CallContext(ctx, &result, "fsn_getLatestNotation", toBlockNumArg(blockNr))
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (ec *Client) GetAddressByNotation(ctx context.Context, notation uint64, blockNr *big.Int) (common.Address, error) {
	var result common.Address
	err := ec.c.CallContext(ctx, &result, "fsn_getAddressByNotation", notation, toBlockNumArg(blockNr))
	if err != nil {
		return common.Address{}, err
	}
	return result, nil
}

func (ec *Client) AllTickets(ctx context.Context, blockNr *big.Int) (map[common.Hash]common.TicketDisplay, error) {
	var result map[common.Hash]common.TicketDisplay
	err := ec.c.CallContext(ctx, &result, "fsn_allTickets", toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ec *Client) AllTicketsByAddress(ctx context.Context, address common.Address, blockNr *big.Int) (map[common.Hash]common.TicketDisplay, error) {
	var result map[common.Hash]common.TicketDisplay
	err := ec.c.CallContext(ctx, &result, "fsn_allTicketsByAddress", address.String(), toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ec *Client) IsAutoBuyTicket(ctx context.Context) (bool, error) {
	var result bool
	err := ec.c.CallContext(ctx, &result, "fsn_isAutoBuyTicket")
	if err != nil {
		return false, err
	}
	return result, nil
}

func (ec *Client) TicketPrice(ctx context.Context, blockNr *big.Int) (*big.Int, error) {
	var result string
	err := ec.c.CallContext(ctx, &result, "fsn_ticketPrice", toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return strToBigInt(result), nil
}

func (ec *Client) TotalNumberOfTickets(ctx context.Context, blockNr *big.Int) (int, error) {
	var result int
	err := ec.c.CallContext(ctx, &result, "fsn_totalNumberOfTickets", toBlockNumArg(blockNr))
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (ec *Client) TotalNumberOfTicketsByAddress(ctx context.Context, address common.Address, blockNr *big.Int) (int, error) {
	var result int
	err := ec.c.CallContext(ctx, &result, "fsn_totalNumberOfTicketsByAddress", address.String(), toBlockNumArg(blockNr))
	if err != nil {
		return 0, err
	}
	return result, nil
}

type RPCAllInfoForAddress struct {
	Tickets   map[common.Hash]common.TicketDisplay `json:"tickets"`
	Balances  map[common.Hash]string               `json:"balances"`
	Timelocks map[common.Hash]*RPCTimeLock         `json:"timeLockBalances"`
	Notation  uint64                               `json:"notation"`
}

func (r *RPCAllInfoForAddress) ToAllInfoForAddress() *common.AllInfoForAddress {
	ret := &common.AllInfoForAddress{
		Tickets:   r.Tickets,
		Balances:  make(map[common.Hash]*big.Int, len(r.Balances)),
		Timelocks: make(map[common.Hash]*common.TimeLock, len(r.Timelocks)),
		Notation:  r.Notation,
	}
	for k, v := range r.Balances {
		ret.Balances[k] = strToBigInt(v)
	}
	for k, v := range r.Timelocks {
		ret.Timelocks[k] = v.ToTimeLock()
	}
	return ret
}

func (ec *Client) AllInfoByAddress(ctx context.Context, address common.Address, blockNr *big.Int) (*common.AllInfoForAddress, error) {
	var result RPCAllInfoForAddress
	err := ec.c.CallContext(ctx, &result, "fsn_allInfoByAddress", address.String(), toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return result.ToAllInfoForAddress(), nil
}

func (ec *Client) GetStakeInfo(ctx context.Context, blockNr *big.Int) (*common.StakeInfo, error) {
	var result common.StakeInfo
	err := ec.c.CallContext(ctx, &result, "fsn_getStakeInfo", toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetBlockReward(ctx context.Context, blockNr *big.Int) (*big.Int, error) {
	var result string
	err := ec.c.CallContext(ctx, &result, "fsn_getBlockReward", toBlockNumArg(blockNr))
	if err != nil {
		return nil, err
	}
	return strToBigInt(result), nil
}

type RPCTransaction struct {
	BlockHash        common.Hash     `json:"blockHash"`
	BlockNumber      *hexutil.Big    `json:"blockNumber"`
	From             common.Address  `json:"from"`
	Gas              hexutil.Uint64  `json:"gas"`
	GasPrice         *hexutil.Big    `json:"gasPrice"`
	Hash             common.Hash     `json:"hash"`
	Input            hexutil.Bytes   `json:"input"`
	Nonce            hexutil.Uint64  `json:"nonce"`
	To               *common.Address `json:"to"`
	TransactionIndex hexutil.Uint    `json:"transactionIndex"`
	Value            *hexutil.Big    `json:"value"`
	V                *hexutil.Big    `json:"v"`
	R                *hexutil.Big    `json:"r"`
	S                *hexutil.Big    `json:"s"`
}

type TxAndReceipt struct {
	FsnTxInput   interface{}            `json:"fsnTxInput,omitempty"`
	Tx           *RPCTransaction        `json:"tx"`
	Receipt      map[string]interface{} `json:"receipt"`
	ReceiptFound bool                   `json:"receiptFound"`
}

func (ec *Client) GetTransactionAndReceipt(ctx context.Context, hash common.Hash) (*TxAndReceipt, error) {
	var result TxAndReceipt
	err := ec.c.CallContext(ctx, &result, "fsn_getTransactionAndReceipt", hash.String())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetTransaction(ctx context.Context, txHash common.Hash) (*RPCTransaction, error) {
	var result RPCTransaction
	err := ec.c.CallContext(ctx, &result, "eth_getTransactionByHash", txHash)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ec *Client) GetTransactionReceipt(ctx context.Context, txHash common.Hash) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := ec.c.CallContext(ctx, &result, "eth_getTransactionReceipt", txHash)
	if err != nil {
		return nil, err
	}
	return result, nil
}
