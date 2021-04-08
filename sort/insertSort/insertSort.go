package insertSort

// insert sort
func InsertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i - 1
		for ; j >= 0 && x < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = x
	}

	return arr
}
