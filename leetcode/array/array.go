package array

import "sort"

/*
题目 #
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:


Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

题目大意 #
给定一个数组，要求在这个数组中找出 3 个数之和为 0 的所有组合。
*/

func Get3Sum(arr []int) (res [][]int) {
	sort.Ints(arr)
	// 双指针
	for index := 1; index < len(arr)-1; index++ {
		// 左指针和右指针
		start, end := 0, len(arr)-1
		// 过滤重复的集合
		if index > 1 && arr[index] == arr[index-1] {
			start = index - 1
		}
		// 双向指针依次找出和为零的组合
		for start < index && index < end {
			// 重复的元素过滤掉
			if start > 1 && arr[start] == arr[start-1] {
				start++
				continue
			}
			// 重复的元素过滤掉
			if end < len(arr)-1 && arr[end] == arr[end-1] {
				end--
				continue
			}
			sum := arr[start] + arr[index] + arr[end]
			if sum == 0 {
				res = append(res, []int{arr[start], arr[index], arr[end]})
				start++
				end--
			} else if sum < 0 {
				start++
			} else {
				end--
			}
		}
	}

	return res
}
