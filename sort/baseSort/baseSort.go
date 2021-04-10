package baseSort

func BaseSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	maxDigit := arrMaxVDigit(arr)
	// 定义每一轮的除数，1,10,100...
	divisor := 1
	// 定义了10个桶，为了防止每一位都一样所以将每个桶的长度设为最大,与原数组大小相同
	bucket := [10][200]int{{0}}
	// 统计每个桶中实际存放的元素个数
	count := [10]int{0}
	// 获取元素中对应位上的数字，即装入那个桶
	var digit int
	// 经过maxDigit+1次装通操作，排序完成
	for i := 1; i <= maxDigit; i++ {
		for j := 0; j < len(arr); j++ {
			tmp := arr[j]
			digit = (tmp / divisor) % 10
			bucket[digit][count[digit]] = tmp
			count[digit]++
		}
		// 被排序数组的下标
		k := 0
		// 从0到9号桶按照顺序取出
		for b := 0; b < 10; b++ {
			if count[b] == 0 {
				continue
			}
			for c := 0; c < count[b]; c++ {
				arr[k] = bucket[b][c]
				k++
			}
			count[b] = 0
		}
		divisor = divisor * 10
	}

	return arr
}

func arrMaxVDigit(arr []int) int {
	// max value
	max := arr[0]
	if len(arr) <= 1 {
		return max
	}

	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	// max value digit
	var maxDigit int
	for max > 0 {
		max /= 10
		maxDigit++
	}

	return maxDigit
}
