# Race detector does not detect data races when using plugins

```
$ go version
go version go1.10 linux/amd64

$ ./test.sh
+ go build -race -o ./plugin/plugin ./plugin
+ + grep -q WARNING: DATA RACE
./plugin/plugin
+ echo As expected, a data race is detected when plugin is running standalone
As expected, a data race is detected when plugin is running standalone
+ go build -race -buildmode=plugin -o ./plugin/plugin.so ./plugin
+ go build -race
+ ./race-nodetect ./plugin/plugin.so
+ grep -q WARNING: DATA RACE
+ echo FAIL: data race NOT detected in plugin
FAIL: data race NOT detected in plugin
```

Reproducible 100% and same behavior on darwin/amd64.
