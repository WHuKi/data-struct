package insertSort

//ShellInsert 希尔排序
func ShellInsert(arr []int, dk int) []int {
	if len(arr) <= 1 {
		return arr
	}

	length := len(arr)
	// 步长
	step := length / dk
	for step >= 1 {
		for i := step; i < length; i++ {
			insertionSortByStep(arr, i)
		}
		step /= dk
	}

	return arr
}

func insertionSortByStep(arr []int, step int) {
	for i := step; i < len(arr); i++ {
		for j := i; j > step-1 && arr[j-step] > arr[j]; j -= step {
			arr[j-step], arr[j] = arr[j], arr[j-step]
		}
	}
}
