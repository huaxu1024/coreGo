package sort

func SortQuick(arr []int, start, end int) {
	if arr == nil || len(arr) < start || start >= end {
		return
	}
	left, right, base := start, end, arr[start]
	for left < right {
		for base <= arr[right] && left < right {
			right --
		}
		arr[left] = arr[right]
		for base > arr[left] && left < right {
			left ++
		}
		arr[right] = arr[left]
	}
	arr[left] = base
	SortQuick(arr, start, left - 1)
	SortQuick(arr, left + 1, end)
}