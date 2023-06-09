package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		dd, _ := strconv.ParseInt(os.Args[1], 10, 32)
		ds, _ := strconv.ParseInt(os.Args[2], 10, 32)
		if dd != 0 && ds != 0 {
			q := dd / ds
			if q > 0 {
				fmt.Printf("résultat: %d\nreste: %d\n", q, dd%ds)
			} else {
				fmt.Println("erreur.")
			}
		} else {
			fmt.Println("erreur.")
		}
	} else {
		fmt.Println("erreur.")
	}
}
