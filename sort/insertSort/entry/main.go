package main

import (
	"data-struct/sort/insertSort"
	"fmt"
)

func main() {
	arr := []int{345, 6, 3}
	newArr := insertSort.ShellInsert(arr, 3)
	fmt.Println(newArr)
}
