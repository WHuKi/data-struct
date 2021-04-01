package main

import (
	"data-struct/queue/loopQueue"
	"fmt"
)

func main() {
	queue := loopQueue.NewLoopQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	if err := queue.EnQueue(33); err != nil {
		fmt.Println(err)
	}

	data1, _ := queue.DeQueue()
	fmt.Println(data1)
	data2, err := queue.DeQueue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data2)
}
