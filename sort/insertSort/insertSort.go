package insertSort

/*
InsertSort insert sort
适用于部分有序，规模比较小的数组
*/
func InsertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}

	return arr
}
