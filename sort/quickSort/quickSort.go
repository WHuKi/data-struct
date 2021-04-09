package quickSort

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	midData := (arr[0] + arr[(len(arr)-1)/2] + arr[len(arr)-1]) / 3
	low, mid, high := make([]int, 0), make([]int, 0), make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] == midData {
			mid = append(mid, arr[i])
		} else if arr[i] < midData {
			low = append(low, arr[i])
		} else {
			high = append(high, arr[i])
		}
	}
	low, high = QuickSort(low), QuickSort(high)

	return append(append(low, mid...), high...)
}
