package main

import (
	"os"
	"time"
)

var counter int

func main() {
	os.Exit(Main())
}

func Main() int {
	go func() {
		counter++
	}()
	time.Sleep(500 * time.Millisecond)
	return counter
}
