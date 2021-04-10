package mergeSort

func MergeSort(arr []int, begin, end int) []int {
	if end-begin <= 1 {
		return arr
	}

	mid := begin + (end-begin+1)/2
	MergeSort(arr, begin, mid)
	MergeSort(arr, mid, end)
	merge(arr, begin, mid, end)

	return arr
}

func merge(arr []int, begin, mid, end int) {
	leftSize := mid - begin
	rightSize := end - mid
	newSize := leftSize + rightSize
	result := make([]int, 0, newSize)

	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lV := arr[begin+l]
		rV := arr[mid+r]
		if lV < rV {
			result = append(result, lV)
			l++
		} else {
			result = append(result, rV)
			r++
		}
	}

	// 将剩下的元素追加到辅助数组后面
	result = append(result, arr[begin+l:mid]...)
	result = append(result, arr[mid+r:end]...)

	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < newSize; i++ {
		arr[begin+i] = result[i]
	}
	return
}
