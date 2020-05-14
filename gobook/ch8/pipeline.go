package ch8

import (
	"fmt"
	"time"
)

func main() {
	time.Tick(1 * time.Second)
}

func goCount() {
	naturals := make(chan int)
	squeares := make(chan int)

	go func() {
		for x := 0; x < 10; x ++  {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squeares <- x * x
		}
		close(squeares)
	}()

	for squeare := range squeares {
		fmt.Println(squeare)
	}
}
