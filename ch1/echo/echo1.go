// demo copyright 2020.03.17 os
package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	seq := " "
	for i := 1; i < len(os.Args); i ++ {
		s += seq + os.Args[i]
	}
	fmt.Println(s)
}
