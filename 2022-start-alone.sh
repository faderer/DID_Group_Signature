#!/bin/bash
nice -50 ./pkg/mod/github.com/go-ethereum-1.10.16/build/bin/geth --datadir ~/.ethereum/net2022-v2 --nodiscover --networkid 2022 console --maxpeers 0 --ws --ws.port 8546 --ws.api "db,eth,net,web3,personl,web3" --http --http.port 8545 --http.addr "0.0.0.0" --http.corsdomain "*" --http.api "eth,net,personal,debug" --allow-insecure-unlock --nousb console
