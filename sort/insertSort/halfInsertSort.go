package insertSort

// half insert sort
func HalfInsertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	var j, low, high int
	for i := 1; i < len(arr); i++ {
		arr[0] = arr[i]
		low, high = 1, i-1
		for low <= high {
			mid := (low + high) / 2
			if arr[0] < arr[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		for j = i - 1; j > high; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = arr[0]
	}

	return arr[1:]
}
