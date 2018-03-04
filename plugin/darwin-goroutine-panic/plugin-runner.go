package main

import (
	"os"
	"plugin"
	"time"
)

func main() {
	p, err := plugin.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	sym, err := p.Lookup("Main")
	if err != nil {
		panic(err)
	}
	f := sym.(func())

	// Note that swapping these two lines fixes the issue.
	go f()
	f()

	time.Sleep(time.Second)
}
