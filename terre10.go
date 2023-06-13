package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		if int(os.Args[1][0]) == int('-') {
			fmt.Println("nombre positif uniquement")
			return
		}
		if strings.ContainsAny(os.Args[1], ".,") {
			fmt.Println("nombre entier uniquement")
			return
		}
		a, _ := strconv.ParseFloat(os.Args[1], 10)
		fmt.Printf("%.0f\n", math.Sqrt(a))
	} else {
		fmt.Println("erreur.")
	}
}
