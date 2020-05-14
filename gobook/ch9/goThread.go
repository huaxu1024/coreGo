package ch9

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func runRun() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

const name string = "abccba"


func testGo() {
	a := []int{1, 2, 3, 4}
	b := make([]int, 0, 4)
	copy(b, a)
	fmt.Printf("%v\n", b)

	b = make([]int, 4, 4)
	copy(b, a)
	fmt.Printf("%v\n", b)
}


