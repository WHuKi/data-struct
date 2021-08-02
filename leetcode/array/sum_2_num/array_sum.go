package sum_2_num

/*
题目 #
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:


Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]

题目大意 #
在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标
*/

func Sum2Num(arr []int) (res [][]int) {
	if len(arr) == 0 {
		return
	}

	vMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if tV, ok := vMap[9-arr[i]]; ok {
			res = append(res, []int{i, tV})
		} else {
			vMap[arr[i]] = i
		}
	}

	return
}

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}
