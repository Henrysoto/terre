package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 3 {
		if int(os.Args[1][0]) == int('-') {
			fmt.Println("exposant négatif")
			return
		}
		a, _ := strconv.ParseFloat(os.Args[1], 10)
		b, _ := strconv.ParseFloat(os.Args[2], 10)
		fmt.Printf("%.0f\n", math.Pow(a, b)) // a**b
	} else {
		fmt.Println("erreur.")
	}
}
