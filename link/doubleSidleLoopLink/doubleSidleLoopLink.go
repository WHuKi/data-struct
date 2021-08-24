package doubleSidleLoopLink

import "fmt"

/*
双向循环链表
*/

type Node struct {
	Pre  *Node
	Next *Node
	Data interface{}
}

//NewNode 初始化
func NewNode() *Node {
	return &Node{}
}

//IsEmpty 判断是否为空
func (node *Node) IsEmpty() bool {
	if node == nil {
		return true
	}

	return false
}

//TraverNode 遍历
func (node *Node) TraverNode() {
	if node == nil {
		return
	}
	head := node

	for node.Next != nil {
		node = node.Next
		fmt.Println(node.Data)
		if node.Next == head {
			break
		}
	}
}

//InsertNodeOfHead 头部插入
func (node *Node) InsertNodeOfHead(data interface{}) *Node {
	if node.IsEmpty() {
		node = NewNode()
	}
	head := node

	newNode := &Node{
		Data: data,
	}
	// 空链表插入
	if head.Next == nil {
		head.Pre = newNode
		head.Next = newNode
		newNode.Pre = head
		// 非空链表插入
	} else {
		head.Next.Pre = newNode
		newNode.Next = head.Next
		head.Next = newNode
		head.Pre = newNode
	}

	return head
}
