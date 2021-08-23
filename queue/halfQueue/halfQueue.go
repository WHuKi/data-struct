package halfQueue

import "sort"

/*
HalfFind
@brief 二分查找
@step:
1. 将数组排序
2. 取中间的值
3. 比较，如果左边的数值比中间值小
*/
func HalfFind(arr []int, num int) (res int, err error) {
	sort.Ints(arr)
	l, r := 0, len(arr)-1
	for i := 0; i < len(arr); i++ {
		midV := (l + r) / 2
		if arr[midV] > num {
			r = midV
		} else if arr[midV] < num {
			l = midV
		} else {
			return midV, nil
		}
	}

	return
}
