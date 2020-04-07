package sort

func SortInsert(arr []int, start, end int) {
	if arr == nil || len(arr) < start || start >= end {
		return
	}
	for i := start; i < end; i++ {
		curr, index := arr[i], i
		for j := i + 1; j < end; j++ {
			if arr[j] < arr[i] {
				arr[i], index = arr[j], j
			}
		}
		arr[index] = curr
	}
}
