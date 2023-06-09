package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) > 1 {
		if unicode.IsLetter(rune(os.Args[1][0])) {
			fmt.Printf("%d\n", len(os.Args[1]))
		} else {
			fmt.Println("erreur.")
		}
	} else {
		fmt.Println("erreur.")
	}
}
