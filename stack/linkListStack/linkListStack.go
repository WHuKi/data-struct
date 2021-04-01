package linkListStack

type Item interface{}

type Node struct {
	Data Item
	Next *Node
}

type Stack struct {
	headNode *Node
}

// push
func (stack *Stack) Push(data Item) {
	newNode := Node{Data: data}
	newNode.Next = stack.headNode
	stack.headNode = &newNode
}

// pop
func (stack *Stack) Pop() Item {
	if stack.headNode == nil {
		return nil
	}
	data := stack.headNode.Data
	stack.headNode = stack.headNode.Next
	return data
}
