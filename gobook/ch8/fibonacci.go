package ch8

import (
	"fmt"
	"time"
)

func fibonacci() {
	start := time.Now()
	const n = 45
	fibN := fib(n)

	fmt.Printf("speed:{%s}, res:{%v}\n", time.Since(start), fibN)
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}

