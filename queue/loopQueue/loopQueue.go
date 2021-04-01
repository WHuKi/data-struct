package loopQueue

import "errors"

var (
	ERR_QUEUE_FULL  = "queue full"
	ERR_QUEUE_EMPTY = "queue empty"
)

type Item interface{}

const QueueSize = 2

type LoopQueue struct {
	Items [QueueSize]Item
	Head  int
	Tail  int
}

func NewLoopQueue() *LoopQueue {
	return &LoopQueue{}
}

// en queue
func (queue *LoopQueue) EnQueue(item Item) (err error) {
	if ((queue.Tail + 1) % QueueSize) == queue.Head {
		return errors.New(ERR_QUEUE_FULL)
	}

	queue.Items[queue.Tail] = item
	queue.Tail = (queue.Tail + 1) % QueueSize
	return
}

// de queue
func (queue *LoopQueue) DeQueue() (item Item, err error) {
	if queue.Head == queue.Tail {
		return nil, errors.New(ERR_QUEUE_EMPTY)
	}

	item = queue.Items[queue.Head]
	queue.Head = (queue.Head + 1) % QueueSize
	return
}
