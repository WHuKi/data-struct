package selectSort

func SelectSort(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		min := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
