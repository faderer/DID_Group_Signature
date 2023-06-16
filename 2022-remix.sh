#!/bin/bash
nice -50 ./pkg/mod/github.com/go-ethereum-1.10.16/build/bin/geth --datadir ~/.ethereum/net2023 --nodiscover --networkid 2022 console --maxpeers 0 --http --http.port 8545 --http.addr "0.0.0.0" --http.corsdomain "https://remix.ethereum.org" --http.api "web,3eth,net,personal,debug" --allow-insecure-unlock --nousb console
