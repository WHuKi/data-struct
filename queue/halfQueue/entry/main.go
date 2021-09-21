package main

import "fmt"

func main() {
	fmt.Println(3 >> 1)

	arr := []int{1, 2, 3, 4, 5}
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + ((r - l) >> 1)
		if arr[mid] == 3 {
			fmt.Println(3, "asdff")
			return
		} else if arr[mid] > 3 {
			r = mid - 1
		} else {
			l = mid + 1
		}
		fmt.Println(mid)
	}

}
