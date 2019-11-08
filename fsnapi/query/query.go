package query

import (
	"math/big"

	"github.com/FusionFoundation/fsn-go-sdk/efsn/ethclient"
)

type QueryClient interface {
	GetAccount(string) (*big.Int, error)
	GetAccountAtBlockNumber(string, int64) (*big.Int, error)
}

type client struct {
	*ethclient.Client
}

func NewClient(url string) (QueryClient, error) {
	ethclient, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return client{ethclient}, nil
}
