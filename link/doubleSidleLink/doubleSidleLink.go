package doubleSidleLink

import "fmt"

/*
@brief 双向链表操作

*/

type Node struct {
	// 前躯节点
	Pre *Node
	// 后继节点
	Next *Node
	// 数据域
	Data interface{}
}

//NewNode 初始化
func NewNode() *Node {
	return &Node{}
}

//IsEmpty 判断节点是否为空
func (node *Node) IsEmpty() bool {
	if node == nil {
		return true
	}

	return false
}

//TraverNode 遍历节点
func (node *Node) TraverNode() {
	if node.IsEmpty() {
		return
	}
	for node.Next != nil {
		node = node.Next
		fmt.Println("节点：", node.Data)
	}
}

//InsertNodeOfTail 尾插入
func (node *Node) InsertNodeOfTail(data interface{}) *Node {
	if node.IsEmpty() {
		node = NewNode()
	}
	temp := node
	// 遍历到最后节点
	for temp.Next != nil {
		temp = temp.Next
	}

	// 最后节点插入元素
	newNode := &Node{
		Data: data,
	}
	temp.Next = newNode
	newNode.Pre = temp

	return node
}

//InsertNodeOfHead 头插入
func (node *Node) InsertNodeOfHead(data interface{}) *Node {
	if node.IsEmpty() {
		node = NewNode()
	}

	temp := node
	newNode := &Node{
		Next: temp.Next,
		Data: data,
	}
	if temp.Next != nil {
		temp.Next.Pre = newNode
	}

	// 处理和头节点的操作
	temp.Next = newNode
	newNode.Pre = temp

	return temp
}

//DelNode 删除节点
func (node *Node) DelNode(data interface{}) {
	if node.IsEmpty() {
		return
	}

	for node.Next != nil {
		node = node.Next
		// 当前节点
		if node.Data.(int) == data.(int) {
			if node.Next != nil {
				node.Next.Pre = node.Pre
				node.Pre.Next = node.Next
				// 删除尾节点
			} else {
				node.Pre.Next = nil
			}
		}
	}
}
