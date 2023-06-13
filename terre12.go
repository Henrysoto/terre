package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		if strings.Contains(os.Args[1], ":") {
			t := strings.Split(os.Args[1], ":")
			hstr := t[0]
			h, _ := strconv.ParseInt(t[0], 10, 32)
			m := t[1]
			if h > 12 {
				fmt.Printf("%d:%sPM\n", h-12, m)
			} else if hstr == "00" || hstr == "0" {
				fmt.Printf("12:%sPM\n", m)
			} else {
				fmt.Printf("%d:%sAM\n", h, m)
			}
		} else {
			fmt.Println("erreur format.")
		}
	} else {
		fmt.Println("erreur.")
	}
}
