package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		if int(os.Args[1][0]) == int('-') {
			fmt.Println("nombre positif uniquement.")
			return
		}
		a, _ := strconv.ParseFloat(os.Args[1], 10)
		if int(a) < 2 {
			fmt.Println("0 ou 1 ne sont pas des nombres premiers.")
			return
		}
		for i := 2; i <= int(math.Sqrt(a)); i++ {
			if int(a)%i == 0 {
				fmt.Printf("Non, %d n'est pas un nombre premier.\n", int(a))
				return
			}
		}
		fmt.Printf("Oui, %d est un nombre premier.\n", int(a))
	} else {
		fmt.Println("erreur.")
	}
}
