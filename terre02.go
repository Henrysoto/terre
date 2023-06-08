package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	_, file, _, ok := runtime.Caller(0)
	if ok {
		s := strings.Split(file, "/")
		fmt.Printf("%s\n", s[len(s)-1])
	}
}
