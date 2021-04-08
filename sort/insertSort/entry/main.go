package main

import (
	"data-struct/sort/insertSort"
	"fmt"
)

func main() {
	arr := []int{1, 4, 3, 5, 4, 6, 2, 3}
	newArr := insertSort.InsertSort(arr)
	fmt.Println(newArr)
}
