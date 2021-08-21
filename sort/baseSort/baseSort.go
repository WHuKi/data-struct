package baseSort

import (
	"math"
	"strconv"
)

/*
version:v1
*/
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

/*
version:v2
*/
//BaseSort 基数排序
func RadixSort(arr []int) []int {
	// 1. 取最大值长度
	maxValueLen := 0
	for i := 0; i < len(arr); i++ {
		// 取数字的最大值并转换为字符串
		n := len(strconv.Itoa(arr[i]))
		if n > maxValueLen {
			maxValueLen = n
		}
	}
	// 2. 排序
	for loc := 1; loc <= maxValueLen; loc++ {
		sort(arr, loc)
	}

	return arr
}

func sort(arr []int, loc int) {
	// 1. 初始化桶
	bucket := make([][]int, 10)
	// 2. 取各个数字到bucket中
	for i := 0; i < len(arr); i++ {
		ji := digit(arr[i], loc)
		if bucket[ji] == nil {
			bucket[ji] = make([]int, 0)
		}
		bucket[ji] = append(bucket[ji], arr[i])
	}
	// 3. 收集二维数组
	for i, idx := 0, 0; i <= 9; i++ {
		for j := 0; j < len(bucket[i]); j++ {
			// 遍历二维数组
			arr[idx] = bucket[i][j] //将数字取出来 给原数组重新赋值
			idx++
		}
	}
}

// 数字，右数第几位，从1开始
func digit(num int, loc int) int {
	return num % int(math.Pow10(loc)) / int(math.Pow10(loc-1))
}
