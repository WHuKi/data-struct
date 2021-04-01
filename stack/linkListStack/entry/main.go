package main

import (
	"data-struct/stack/linkListStack"
	"fmt"
)

func main() {
	stack := linkListStack.Stack{}
	stack.Push(1)
	data := stack.Pop()
	fmt.Println(data)
}
