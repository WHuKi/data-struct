package quickSort

//QuickSort 普通快速排序
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	midData := (arr[0] + arr[(len(arr)-1)/2] + arr[len(arr)-1]) / 3
	low, mid, high := make([]int, 0), make([]int, 0), make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] == midData {
			mid = append(mid, arr[i])
		} else if arr[i] < midData {
			low = append(low, arr[i])
		} else {
			high = append(high, arr[i])
		}
	}
	low, high = QuickSort(low), QuickSort(high)

	return append(append(low, mid...), high...)
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
