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

package mongodb

const (
	tbSyncInfo     string = "SyncInfo"
	tbBlocks       string = "Blocks"
	tbTransactions string = "Transactions"
	tbContracts    string = "Contracts"

	KeyOfLatestSyncInfo string = "latest"
)

type MgoSyncInfo struct {
	Key             string `bson:"_id"`
	Number          uint64 `bson:"number"`
	Timestamp       uint64 `bson:"timestamp"`
	TotalDifficulty uint64 `bson:"totalDifficulty"`
}

type MgoBlock struct {
	Key             string   `bson:"_id"`
	Number          uint64   `bson:"number"`
	Hash            string   `bson:"hash"`
	ParentHash      string   `bson:"parentHash"`
	Nonce           uint64   `bson:"ticketOrder"` // spec
	Miner           string   `bson:"miner"`
	Difficulty      uint64   `bson:"difficulty"`
	TotalDifficulty uint64   `bson:"totalDifficulty"`
	Size            uint64   `bson:"size"`
	GasLimit        uint64   `bson:"gasLimit"`
	GasUsed         uint64   `bson:"gasUsed"`
	Timestamp       uint64   `bson:"timestamp"`
	BlockTime       uint64   `bson:"blockTime"`
	TxCount         int      `bson:"txcount"`
	AvgGasprice     string   `bson:"avgGasprice"`
	Reward          string   `bson:"reward"`
	SelectedTicket  string   `bson:"selectedTicket"` // spec
	RetreatTickets  []string `bson:"retreatTickets"` // spec
	RetreatMiners   []string `bson:"retreatMiners"`  // spec
	TicketNumber    int      `bson:"ticketNumber"`   // spec
}

type MgoTransaction struct {
	Key              string      `bson:"_id"`
	Hash             string      `bson:"hash"`
	Nonce            uint64      `bson:"nonce"`
	BlockHash        string      `bson:"blockHash"`
	BlockNumber      uint64      `bson:"blockNumber"`
	TransactionIndex int         `bson:"transactionIndex"`
	From             string      `bson:"from"`
	To               string      `bson:"to"`
	Value            string      `bson:"value"`
	ValueInt         uint64      `bson:"ivalue"`
	ValueDec         uint64      `bson:"dvalue"`
	GasLimit         uint64      `bson:"gasLimit"`
	GasPrice         string      `bson:"gasPrice"`
	GasUsed          uint64      `bson:"gasUsed"`
	Timestamp        uint64      `bson:"timestamp"`
	Input            string      `bson:"input"`
	Status           uint64      `bson:"status"`
	CoinType         string      `bson:"coinType"`
	Type             string      `bson:"type"` // spec
	Log              interface{} `bson:"log"`  // spec

	Erc20Receipts    []*Erc20Receipt    `bson:"erc20Receipts,omitempty"`
	ExchangeReceipts []*ExchangeReceipt `bson:"exchangeReceipts,omitempty"`
}

// Erc20Receipt erc20 tx receipt
type Erc20Receipt struct {
	LogType  string `bson:"logType"`
	LogIndex int    `bson:"logIndex"`
	Erc20    string `bson:"erc20"`
	From     string `bson:"from"`
	To       string `bson:"to"`
	Value    string `bson:"value"`
}

// ExchangeReceipt exchange tx receipt
type ExchangeReceipt struct {
	LogType         string `bson:"txnsType"`
	LogIndex        int    `bson:"logIndex"`
	Exchange        string `bson:"exchange"`
	Address         string `bson:"address"`
	TokenFromAmount string `bson:"tokenFromAmount"`
	TokenToAmount   string `bson:"tokenToAmount"`
}

type MgoContract struct {
	Key  string `bson:"_id"`
	Type string `bson:"type"` // ERC20, MBTC, ...
}
