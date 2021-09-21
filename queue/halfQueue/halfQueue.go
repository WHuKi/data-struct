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

/*
二分查找变形
1. 查找第一个等于给定值的元素
if mid -1 >=0 {
	if mid ==0 || arr[mid -1] != value {return mid}else {
		high = mid -1
	}
}

2. 查找最后一个等于给定值的元素
if mid +1 < n {
	if (mid == n - 1) || arr[mid + 1] != value {return mid}else {
		low = mid + 1
	}
}

3. 查找第一个大于等于给定值的元素
if (a[mid] >= value) {
	if mid ==0 || a[mid -1 ] < value {return mid}else{
		high = mid+1
 	}
}

4. 查找最后一个小于等于给定值的元素
*/
