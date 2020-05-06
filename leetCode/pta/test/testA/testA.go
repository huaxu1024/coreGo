package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func maxin() {

	reader := bufio.NewReader(os.Stdin)
	n := readInt(reader)

	for i := 0; i < n; i++ {
		m := readInt(reader)
		arr := readArray(reader)

		for {
			sort.Ints(arr)
			if arr[0] == arr[m-1] {
				fmt.Println(arr[0])
				break
			}
			val := arr[m-1] - arr[0]
			arr[0] = val
			arr[m-1] = val
		}
	}
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimRight(line, "\n")
}

func readInt(reader *bufio.Reader) int {
	num, _ := strconv.Atoi(readLine(reader))
	return num
}

func readArray(reader *bufio.Reader) []int {
	line := readLine(reader)
	strs := strings.Split(line, " ")
	nums := make([]int, len(strs))
	for i, s := range strs {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}