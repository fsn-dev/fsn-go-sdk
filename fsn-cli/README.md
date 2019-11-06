# fsn-cli Usage

## Introduction

`fsn-cli` is mainly used to offlinely create transaction and sign transaction,
and also support query balance or other chain status onlinely through RPC calling.

## Get help info

```
./bin/fsn-cli --help # get supported commands
./bin/fsn-cli buyticket --help # get command help info
```

## For developers

If you want to contribute to this project, you can verify your result in the following way.

take `buyticket` for example,

1. create a raw transaction
```
./bin/fsn-cli buyticket --start 1573019189 --end 1575611189 --sign --from 0x37a200388caa75edcc53a2bd329f7e9563c6acb6 --nonce 100 --keystore mykeystorefile --password mypasswordfile 

0xf87464843b9aca0083015f9094ffffffffffffffffffffffffffffffffffffffff808ecd048bca845dc25e35845de9eb3582ff49a0efdf19823d1bcef034c0684fb1ce6686d5e940e84573f102cecb9b9baeb4f841a042fd501347a89038a7deb390bcd24c3261e0bfaa8d76f04f114902bbab73075e
```

2. decode and check the raw transaction
```
./bin/fsn-cli decoderawtx 0xf87464843b9aca0083015f9094ffffffffffffffffffffffffffffffffffffffff808ecd048bca845dc25e35845de9eb3582ff49a0efdf19823d1bcef034c0684fb1ce6686d5e940e84573f102cecb9b9baeb4f841a042fd501347a89038a7deb390bcd24c3261e0bfaa8d76f04f114902bbab73075e | jq

{
  "nonce": "0x64",
  "gasPrice": "0x3b9aca00",
  "gas": "0x15f90",
  "to": "0xffffffffffffffffffffffffffffffffffffffff",
  "value": "0x0",
  "input": "0xcd048bca845dc25e35845de9eb35",
  "v": "0xff49",
  "r": "0xefdf19823d1bcef034c0684fb1ce6686d5e940e84573f102cecb9b9baeb4f841",
  "s": "0x42fd501347a89038a7deb390bcd24c3261e0bfaa8d76f04f114902bbab73075e",
  "hash": "0xcc6a902c85edd796702ebd0cbefe5ae4416dfa0763a47c2b02e6c323dca9d68f"
}
```

3. decode and check the raw transaction's input data
```
./bin/fsn-cli decoderawtx --input 0xcd048bca845dc25e35845de9eb35 | jq

{
  "FuncType": "BuyTicketFunc",
  "FuncParam": {
    "Start": 1573019189,
    "End": 1575611189
  }
}
```
