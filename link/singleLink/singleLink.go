package singleLink

import "fmt"

/*
@brief 单链表操作
func
1. 创建
2. 删除
3. 修改
4. 头、尾插入
*/

type Node struct {
	// 指针域
	NextPtr *Node
	// 数据域
	Data interface{}
}

//CreateHeadNode 创建头节点
func CreateHeadNode() *Node {
	return &Node{}
}

//InsertNode 尾插入
func (node *Node) InsertNode(data interface{}) *Node {
	// 判断是否为空
	if node.IsEmpty() {
		node = CreateHeadNode()
	}
	// 走到最后的节点
	for node.NextPtr != nil {
		node = node.NextPtr
	}

	node.NextPtr = &Node{
		Data: data,
	}

	return node
}

//InsertNodeOfHead 头节点插入
func (node *Node) InsertNodeOfHead(data interface{}) *Node {
	// 判断是否为空
	if node.IsEmpty() {
		node = CreateHeadNode()
	}
	// 新节点
	newNode := &Node{
		Data: data,
	}

	temNode := node.NextPtr
	// 链表中有元素
	newNode.NextPtr = temNode
	node.NextPtr = newNode

	return node
}

//TraverseLink 遍历链表
func (node *Node) TraverseLink() {
	if node.IsEmpty() {
		return
	}
	for node.NextPtr != nil {
		node = node.NextPtr
		fmt.Println(node.Data)
	}
}

//DelLinkNode 删除元素
func (node *Node) DelLinkNode(data interface{}) {
	if node.IsEmpty() {
		return
	}

	for node.NextPtr != nil {
		if node.NextPtr.Data.(int) == data.(int) {
			node.NextPtr = node.NextPtr.NextPtr
		}

		node = node.NextPtr
	}
}

//IsEmpty 判断链表是否为空
func (node *Node) IsEmpty() bool {
	if node == nil {
		return true
	}
	return false
}
