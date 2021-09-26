package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 40, 5, 30, 3, 40, 3, 50, 3, 0, 4, 3, 2, 45, 56, 78, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 550}

	//fmt.Println(heapSort.NMaxV(arr, 5))

	HeapSort(arr)
	fmt.Println(arr)
}

//HeapSort 堆排序
func HeapSort(arr []int) {
	var n = len(arr) - 1
	for root := n / 2; root >= 0; root-- {
		minHeap(root, n, arr)
	}

	for end := n; end >= 0; end-- {
		if arr[0] < arr[end] {
			arr[0], arr[end] = arr[end], arr[0]
			minHeap(0, end-1, arr)
		}
	}

}

func minHeap(root, end int, c []int) {
	for {
		// 子节点
		var child = 2*root + 1
		// 和最后的节点比较
		if child > end {
			break
		}

		// 同级节点比较 （找出左右节点中最大的那个）
		if child+1 < end && c[child] > c[child+1] {
			child += 1
		}

		// 数据交换
		if c[root] > c[child] {
			c[root], c[child] = c[child], c[root]
			root = child
		} else {
			break
		}

	}
}
