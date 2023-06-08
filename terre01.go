package main

import (
	"fmt"
)

func main() {
	for i := uint8('a'); i <= uint8('z'); i++ {
		fmt.Print(string(i))
	}
	fmt.Println()
}
