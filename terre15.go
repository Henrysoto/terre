package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) >= 4 {
		t := make([]int64, 0, len(os.Args[1:]))
		for _, n := range os.Args[1:] {
			x, _ := strconv.ParseInt(n, 10, 32)
			t = append(t, x)
		}
		srt := true
		for i := 0; i <= len(t); i++ {
			if i+1 >= len(t) {
				break
			}
			if t[i] > t[i+1] {
				srt = false
			}
		}
		if srt {
			fmt.Println("Triée !")
		} else {
			fmt.Println("Pas triée !")
		}
	} else {
		fmt.Println("erreur.")
	}
}
