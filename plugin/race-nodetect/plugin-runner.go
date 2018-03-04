package main

import (
	"os"
	"plugin"
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
	f := sym.(func() int)
	_ = f()
}
