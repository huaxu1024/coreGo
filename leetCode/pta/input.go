package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 2000*1024), 2000)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	for i := 0; i < N; i ++ {
		sc.Scan()
		arr := strings.Split(sc.Text(), " ")
		fmt.Printf("%v", arr)
	}
}

