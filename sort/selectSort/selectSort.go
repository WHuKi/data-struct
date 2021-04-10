package selectSort

func SelectSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		k := i
		for j := i + 1; j < len(arr); j++ {
			if arr[k] > arr[j] {
				k = j
			}
		}
		if k != i {
			arr[i], arr[k] = arr[k], arr[i]
		}
	}
	return arr
}
