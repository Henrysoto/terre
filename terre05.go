package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if int(arg[0]) == int('-') {
			fmt.Println("Nombre négatif")
			return
		}
		if unicode.IsDigit(rune(arg[0])) {
			nb, _ := strconv.Atoi(arg)
			if nb%2 == 0 {
				fmt.Println("pair")
			} else {
				fmt.Println("impair")
			}
		} else {
			fmt.Println("Tu ne me la mettras pas à l'envers.")
		}
	} else {
		fmt.Println("Tu ne me la mettras pas à l'envers.")
	}
}
