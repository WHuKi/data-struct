package main

import (
	"data-struct/sort/insertSort"
	"fmt"
)

func main() {
	arr := []int{1, 4, 7, 20, 4}
	newArr := insertSort.ComoSort(arr)
	fmt.Println(newArr)
}
