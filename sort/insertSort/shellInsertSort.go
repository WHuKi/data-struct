package insertSort

// shell insert sort
func ShellInsert(arr []int, dk int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for i := dk + 1; i < len(arr); i++ {
		if arr[i] < arr[i-dk] {
			arr[0] = arr[i]
			j := i - dk
			for ; j > 0 && arr[0] < arr[j]; j = j - dk {
				arr[j+dk] = arr[j]
			}
			arr[j+dk] = arr[0]
		}
	}

	return arr[1:]
}
