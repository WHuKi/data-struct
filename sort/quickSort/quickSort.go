package quickSort

//QuickSort 普通快速排序
/*
@brief: 快速排序
@step:
1. 找出基准值
2. 定义左向量、右向量
	2.1 并且左向量向右移动直到找到比基准值大的数值
	2.2 右向量向左移动直到找到比基准值小的数值
	2.3 如果此时左向量下标比右向量的下标小则交换左向量和右向量的值，左向量和右向量都加一
3. 如果左向量比尾节点的数值小则递归该函数
4. 如果右向量比左节点的数值小则递归该函数
*/
func QuickSort(startIndex, endIndex int, arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 1. 定义基准值、左向量、右向量
	pivot := arr[(startIndex+endIndex)/2]
	l, r := startIndex, endIndex
	// 2. 处理
	for l <= r {
		// 左向量
		for arr[l] < pivot {
			l++
		}
		// 右向量
		for arr[r] > pivot {
			r--
		}
		// 处理
		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	// 3. 如果左向量比尾节点的数值小则递归该函数
	if l < endIndex {
		QuickSort(l, endIndex, arr)
	}
	// 4. 如果右向量比左节点的数值小则递归该函数
	if r > startIndex {
		QuickSort(startIndex, r, arr)
	}

	return arr
}

//ThreeWayQuickSort 三向切分快速排序
func ThreeWayQuickSort(s []int) {
	sort3way(s, 0, len(s)-1)
}

//在lt之前的(lo~lt-1)都小于中间值
//在gt之前的(gt+1~hi)都大于中间值
//在lt~i-1的都等于中间值
//在i~gt的都还不确定（最终i会大于gt，即不确定的将不复存在）
func sort3way(s []int, lo, hi int) {
	if lo >= hi {
		return
	}
	v, lt, i, gt := s[lo], lo, lo+1, hi
	for i <= gt {
		if s[i] < v {
			swap(s, i, lt)
			lt++
			i++
		} else if s[i] > v {
			swap(s, i, gt)
			gt--
		} else {
			i++
		}
	}
	sort3way(s, lo, lt-1)
	sort3way(s, gt+1, hi)
}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}
