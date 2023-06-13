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
			h, _ := strconv.ParseInt(t[0], 10, 32)
			mstr := t[1]
			var m string
			var af bool
			if strings.Contains(mstr, "A") {
				m = mstr[:strings.Index(mstr, "A")]
				af = false
			} else if strings.Contains(mstr, "P") {
				m = mstr[:strings.Index(mstr, "P")]
				af = true
			} else {
				fmt.Println("erreur format.")
				return
			}
			if af {
				if h == 12 {
					fmt.Printf("00:%s\n", m)
				} else {
					fmt.Printf("%d:%s\n", h+12, m)
				}
			} else {
				fmt.Printf("%d:%s\n", h, m)
			}
		} else {
			fmt.Println("erreur format.")
		}
	} else {
		fmt.Println("erreur.")
	}
}
