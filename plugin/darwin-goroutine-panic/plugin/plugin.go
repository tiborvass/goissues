package main

import (
	"fmt"
	"runtime"
)

func Main() {
	fmt.Println("[plugin] NumGoroutine:", runtime.NumGoroutine())
}
