#!/bin/bash
#rm -rf ~/.ethereum/net2022
./pkg/mod/github.com/go-ethereum-1.10.16/build/bin/geth --datadir ~/.ethereum/net2022-v2 --networkid 2022 init pkg/mod/github.com/go-ethereum-1.10.16/build/bin/privateNetwork/genesis.json
