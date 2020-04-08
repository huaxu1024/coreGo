package dfs

import "testing"

func TestDfs(t *testing.T) {
	Dfs(0, 3, []int{1,2,3}, make([]bool, 3))
}