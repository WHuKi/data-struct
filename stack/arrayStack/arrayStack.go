package arrayStack

import "errors"

var (
	ERR_STACK_FULL = "stack is full"
)

type Item interface{}

type ItemStack struct {
	Items []Item
	N     int
}

// new stack
func NewStack() *ItemStack {
	return &ItemStack{}
}

// push
func (stack *ItemStack) Push(item Item) (err error) {
	if len(stack.Items) > stack.N {
		return errors.New(ERR_STACK_FULL)
	}

	stack.Items = append(stack.Items, item)
	return nil
}

// pop
func (stack *ItemStack) Pop() Item {
	if len(stack.Items) == 0 {
		return nil
	}

	item := stack.Items[len(stack.Items)-1]
	stack.Items = stack.Items[0 : len(stack.Items)-1]
	return item
}
