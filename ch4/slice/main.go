package main

import "fmt"

func reverse() {
	s := []int{0,1,2,3,4,5,6,7}
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}

func appendDD() {
	var runes []rune
	for _, r := range "hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < len(x) * 2 {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	s := []int{1,2,3,4,5,6}
	fmt.Println(remove(s, 2))
}
