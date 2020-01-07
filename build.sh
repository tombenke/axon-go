#!/bin/bash

mkdir -p ./dist/linux-386/
#mkdir -p ./dist/linux-amd64/
mkdir -p ./dist/linux-arm5/
#mkdir -p ./dist/windows-386/
#mkdir -p ./dist/windows-amd64/
#mkdir -p ./dist/darwin-386/
#mkdir -p ./dist/darwin-amd64/

env GOOS=linux GOARCH=386 go build -o dist/linux-386/  github.com/tombenke/axon-go/axon-cron/
env GOOS=linux GOARCH=386 go build -o dist/linux-386/  github.com/tombenke/axon-go/axon-debug/
env GOOS=linux GOARCH=386 go build -o dist/linux-386/  github.com/tombenke/axon-go/axon-influxdb-writer/

env GOOS=linux GOARCH=arm GOARM=5 go build -o dist/linux-arm5/  github.com/tombenke/axon-go/axon-cron/
env GOOS=linux GOARCH=arm GOARM=5 go build -o dist/linux-arm5/  github.com/tombenke/axon-go/axon-debug/
env GOOS=linux GOARCH=arm GOARM=5 go build -o dist/linux-arm5/  github.com/tombenke/axon-go/axon-influxdb-writer/

#env GOOS=windows GOARCH=amd64 go build -o dist/  github.com/tombenke/axon-go/axon-cron/
