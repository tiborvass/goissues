#!/bin/sh

set -x

go build -o plugin/plugin.so -buildmode=plugin ./plugin
go build
./darwin-goroutine-panic ./plugin/plugin.so
