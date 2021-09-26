package mpSort

// MpSort 冒泡排序
func MpSort(arr []int) (res []int) {
	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
