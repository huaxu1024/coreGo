package sort

func SortBubbe(arr []int, start, end int) []int {
	if end <= start || len(arr) < start {
		return nil
	}
	for i := start; i < end; i ++ {
		for j := i + 1; j < end; j ++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
