package linkListQueue

import "errors"

var (
	ERR_QUEUE_EMPTY = "queue empty"
)

type Item interface{}

type Node struct {
	Data Item
	Next *Node
}

type Queue struct {
	headNode *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

// en queue
func (queue *Queue) Enqueue(data Item) {
	node := &Node{Data: data}
	if queue.headNode == nil {
		queue.headNode = node
	} else {
		currentNode := queue.headNode
		if currentNode.Next != nil {
			currentNode = currentNode.Next
		}

		currentNode.Next = node
	}
}

// de queue
func (queue *Queue) DeQueue() (item Item, err error) {
	if queue.headNode == nil {
		return nil, errors.New(ERR_QUEUE_EMPTY)
	}

	item = queue.headNode.Data
	queue.headNode = queue.headNode.Next
	return
}
