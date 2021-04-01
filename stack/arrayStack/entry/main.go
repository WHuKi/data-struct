package main

import (
	"data-struct/stack/arrayStack"
	"fmt"
)

func main() {
	stack := arrayStack.NewStack()
	if err := stack.Push(1); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(stack.Pop())
}
