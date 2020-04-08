package dfs

import "fmt"

func Dfs(step, length int, a []int, book []bool) {
	if step == length {
		for _, val := range a {
			fmt.Print(" ", val)
		}
		fmt.Println()
		return
	}

	for i := 0; i < length; i++ {
		if !book[i] {
			a[step] = i + 1
			book[i] = true
			Dfs(step + 1, length, a, book)
			book[i] = false
		}
	}
}