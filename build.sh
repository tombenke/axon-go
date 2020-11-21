#!/bin/bash

rm -fr ./dist/*

mkdir -p ./dist/linux-386/
mkdir -p ./dist/linux-amd64/
mkdir -p ./dist/linux-arm5/
mkdir -p ./dist/windows-386/
mkdir -p ./dist/windows-amd64/
mkdir -p ./dist/darwin-386/
mkdir -p ./dist/darwin-amd64/

build () {

    if [ -z ${5} ]; then
        env GOOS=${3} GOARCH=${4} go build -o dist/${2}/ github.com/tombenke/axon-go/${1}/
    else
        env GOOS=${3} GOARCH=${4} GOARM=${5} go build -o dist/${2}/ github.com/tombenke/axon-go/${1}/
    fi
}

agents="axon-cron axon-debug axon-file-reader axon-influxdb-writer axon-function-js"

for agent in $agents
do
    echo Build "$agent..."
    build "$agent" "linux-386" "linux" "386"
    build "$agent" "linux-amd64" "linux" "amd64"
    build "$agent" "linux-arm5" "linux" "arm" "5"
    build "$agent" "windows-386" "windows" "386"
    build "$agent" "windows-amd64" "windows" "amd64"
    build "$agent" "darwin-386" "darwin" "386"
    build "$agent" "darwin-amd64" "darwin" "amd64"
    echo
done

echo "Build complete."

