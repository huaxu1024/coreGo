package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	v := 1
	incr(&v)
	fmt.Println(&v)
	fmt.Println(incr(&v))
	fmt.Println(&v)

	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func incr(i *int) int{
	*i ++
	return *i
}


func fuzhi() {
	medals := []string{"a", "b", "c"}
	fmt.Println(medals)
}