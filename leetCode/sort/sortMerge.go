package sort

func SortMerge(arr []int, lo, hi int) {
	if arr == nil || len(arr) < lo || lo >= hi {
		return
	}
	mid := lo + (hi - lo) / 2
	SortMerge(arr, lo, mid)
	SortMerge(arr, mid + 1, hi)
	merge(arr, lo, mid, hi)
}

func merge(array []int, lo, mid, hi int) {
	index, left, right, length := 0, lo, mid + 1, hi - lo + 1
	copy := make([]int, length)

	for index < length {
		if left > mid {
			copy[index] = array[right]
			right ++

		} else if right > hi {
			copy[index] = array[left]
			left ++

		} else if array[left] <= array[right] {
			copy[index] = array[left]
			left ++

		} else {
			copy[index] = array[right]
			right ++
		}
		index ++
	}
	for i := 0; i < length; i ++ {
		array[i + lo] = copy[i]
	}
}
