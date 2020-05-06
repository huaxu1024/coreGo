package main

import (
"bufio"
"fmt"
"os"
"strconv"
"strings"
)

func main() {

	arr := make([]int64, 50)
	arr[1] = 2
	for i := 2; i < 50; i ++ {
		arr[i] = arr[i - 1] * 2
	}

	reader := bufio.NewReader(os.Stdin)
	n := readInt(reader)

	for j := 0; j < n; j ++ {
		num := int64(readInt(reader))
		if num < 4 {
			fmt.Println(1)
		} else {
			for i := 49; i >= 1; i -- {
				if num >= arr[i] {
					fmt.Println(i)
					break
				}
			}
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

func readUint(reader *bufio.Reader) int64 {
	num, _ := strconv.ParseInt(readLine(reader),10, 64)
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