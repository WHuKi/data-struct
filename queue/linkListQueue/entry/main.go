package main

import (
	"data-struct/queue/linkListQueue"
	"fmt"
)

func main() {
	queue := linkListQueue.NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	fmt.Println(queue.DeQueue())
}
