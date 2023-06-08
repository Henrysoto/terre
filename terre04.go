package main

import (
	"fmt"
	"os"
)

func main() {
	for i := uint8(rune(os.Args[1][0])); i <= uint8('z'); i++ {
		fmt.Print(string(i))
	}
	fmt.Println()
}
