package sort

func SortHeap(arr []int, start, end int) *[]int {
	if arr == nil || len(arr) < start || start >= end {
		return nil
	}
	for i := (end >> 1) + 1; i >= start; i-- {
		buildHead(arr, i, end)
	}
	for i := end - 1; i > start; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		buildHead(arr, 0, i)
	}
	return &arr
}

func buildHead(arr []int, index, end int) {
	base := arr[index]
	for i := (index << 1) + 1; i < end; i = (index << 1) + 1 {
		if i+1 < end && arr[i] < arr[i+1] {
			i++
		}
		if base > arr[i] {
			break
		}
		arr[index] = arr[i]
		index = i
	}
	arr[index] = base
}
