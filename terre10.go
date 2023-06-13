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
			fmt.Println("exposant négatif")
			return
		}
		a, _ := strconv.ParseFloat(os.Args[1], 10)
		fmt.Printf("%.0f\n", math.Sqrt(a))
	} else {
		fmt.Println("erreur.")
	}
}
