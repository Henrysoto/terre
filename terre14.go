package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) == 4 {
		t := make([]int64, 0, len(os.Args[1:]))
		for _, n := range os.Args[1:] {
			x, _ := strconv.ParseInt(n, 10, 32)
			t = append(t, x)
		}
		sort.Slice(t, func(i, j int) bool {
			return t[i] < t[j]
		})
		fmt.Printf("%d\n", t[1])
	} else {
		fmt.Println("erreur.")
	}
}
