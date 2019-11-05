# mongosync usage

## Introduction

`mongosync` is a fusion chain block syncer.
It will connect to a fusion node through "http", "https", "ws", "wss", or IPC file protocal. After connected, it will retreive blocks in specified range through RPC calling, and store results of blocks and transactions information into local mongodb database.

this program use multi-goroutines to syncing blocks, it can sync 800,000 blocks in less than 15 minutes.

## prerequisite

1. install mongodb
2. start mongod service
3. start fusion node (--gcmode archive)

## get help info

* ./bin/mongosync --help
```
NAME:
   mongosync - sync blocks into mongodb

USAGE:
   mongosync [global options] command [command options] [arguments...]

VERSION:
   1.0.0-10fd3ad5

COMMANDS:
     syncAll  sync all info into mongodb
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbosity value  log verbosity (0-9) (default: 3)
   --jobs value       number of jobs (1-1000) (default: 10)
   --help, -h         show help
   --version, -v      print the version
```

`--verbosity` specify log level to show, won't show log info with level greater than this value.

log level from 0 to 5 are: CRIT, ERROR, WARN,  INFO, DEBUG, TRACE

`--jobs` specify how many workers(goroutines) to the syncing job.

* ./bin/mongosync syncAll --help
```
NAME:
   mongosync syncAll - sync all info into mongodb

USAGE:
   mongosync syncAll [command options] <severURL>

DESCRIPTION:

Sync all info into mongodb.
severURL support "http", "https", "ws", "wss", "stdio", IPC file

OPTIONS:
   --mongoURL value  mongodb URL (default: "localhost:27017")
   --dbName value    database name (default: "fusion")
   --stable value    sync blocks that is stable lower than the latest (default: 10)
   --start value     sync blocks from this height (default: 0)
   --end value       sync blocks to this height (not inclusive) (default: 0)
   --overwrite       overwrite synced entries
```

when restart `mongosync` without specify `--start` or start is 0, then will resync from `SyncInfo.Number`.

if `--end` is not specified or end is 0, then will loop sync the latest blocks forever.

if `--overwrite` is specified, then will overwrite document (record) with same \_id in syncing, otherwise it will ignore that document and continue.

## examples

```
./mongosync --jobs=20 syncAll --dbName=fusion ipcfile

./mongosync --jobs=20 syncAll --dbName=fusion --mongoURL=username:password@localhost:27017 ipcfile
```

## database info

it will create 3 collections (tables): `SyncInfo`, `Blocks`, `Transactions`

### SyncInfo

| Name | Type | Key |
| ------ | ------ | ------ |
|Number          | uint64 | `bson:"number"` |
|Timestamp       | uint64 | `bson:"timestamp"` |
|TotalDifficulty | uint64 |`bson:"totalDifficulty"` |

### Blocks

| Name | Type | Key |
| ------ | ------ | ------ |
|Number            |uint64   |`bson:"number"`|
|Hash              |string   |`bson:"hash"`|
|ParentHash        |string   |`bson:"parentHash"`|
|Nonce             |uint64   |`bson:"ticketOrder"` // spec|
|Miner             |string   |`bson:"miner"`|
|Difficulty        |uint64   |`bson:"difficulty"`|
|TotalDifficulty   |uint64   |`bson:"totalDifficulty"`|
|Size              |uint64   |`bson:"size"`|
|GasLimit          |uint64   |`bson:"gasLimit"`|
|GasUsed           |uint64   |`bson:"gasUsed"`|
|Timestamp         |uint64   |`bson:"timestamp"`|
|BlockTime         |uint64   |`bson:"blockTime"`|
|TxCount           |int      |`bson:"txcount"`|
|AvgGasprice       |string   |`bson:"avgGasprice"`|
|Reward            |string   |`bson:"reward"`|
|SelectedTicket    |string   |`bson:"selectedTicket"` // spec|
|RetreatTickets    |[]string |`bson:"retreatTickets"` // spec|
|RetreatMiners     |[]string |`bson:"retreatMiners"`  // spec|
|TicketNumber      |int      |`bson:"ticketNumber"`   // spec|

### Transactions

| Name | Type | Key |
| ------ | ------ | ------ |
|Hash             |string      |`bson:"hash"`|
|Nonce            |uint64      |`bson:"nonce"`|
|BlockHash        |string      |`bson:"blockHash"`|
|BlockNumber      |uint64      |`bson:"blockNumber"`|
|TransactionIndex |int         |`bson:"transactionIndex"`|
|From             |string      |`bson:"from"`|
|To               |string      |`bson:"to"`|
|Value            |string      |`bson:"value"`|
|ValueInt         |uint64      |`bson:"ivalue"`|
|ValueDec         |uint64      |`bson:"dvalue"`|
|GasLimit         |uint64      |`bson:"gasLimit"`|
|GasPrice         |string      |`bson:"gasPrice"`|
|GasUsed          |uint64      |`bson:"gasUsed"`|
|Timestamp        |uint64      |`bson:"timestamp"`|
|Input            |string      |`bson:"input"`|
|Status           |uint64      |`bson:"status"`|
|CoinType         |string      |`bson:"coinType"`|
|Type             |string      |`bson:"type"` // spec|
|Log              |interface{} |`bson:"log"`  // spec|

