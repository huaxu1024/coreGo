package testB

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	num := readInt(reader)

	for i := 0; i < num; i++ {
		nums := readArray(reader)
		arr := readArray(reader)
		sort.Sort(sort.IntSlice(arr))

		var cnt = uint64(0)
		for j, m2_1 := 0, nums[1] * 2 - 1; j < nums[1]; j ++ {
			cnt += uint64(arr[j] * arr[m2_1 - j])
		}
		fmt.Println(cnt)
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