package main

import (
	"fmt"
	"time"
)

func test() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Println(fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "abcdef" {
			fmt.Println(r)
			time.Sleep(delay)
		}
	}

}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}

