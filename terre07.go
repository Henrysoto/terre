package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		var p string = ""
		for i := 1; i <= len(os.Args)-1; i++ {
			for k := len(os.Args[i]) - 1; k >= 0; k-- {
				p = fmt.Sprintf("%s%c", p, rune(os.Args[i][k]))
			}
		}
		fmt.Println(p)
	} else {
		fmt.Println("erreur.")
	}
}
