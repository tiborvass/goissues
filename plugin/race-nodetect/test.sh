#!/bin/sh

set -x

go build -race -o ./plugin/plugin ./plugin
if 2>&1 ./plugin/plugin | grep -q 'WARNING: DATA RACE'; then
	echo "As expected, a data race is detected when plugin is running standalone"
else
	echo "Expecting to detect a data race when plugin is running standalone"
	exit 1
fi
go build -race -buildmode=plugin -o ./plugin/plugin.so ./plugin

go build -race
if 2>&1 ./race-nodetect ./plugin/plugin.so | grep -q 'WARNING: DATA RACE'; then
	echo "PASS: data race detected in plugin"
else
	echo "FAIL: data race NOT detected in plugin"
fi

