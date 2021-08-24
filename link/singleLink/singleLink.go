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
	for node != nil {
		fmt.Println(node.Data)
		node = node.NextPtr
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

//GetMidNode 找到中间节点
/*
@brief
1. 利用快慢指针的思路来获取中间节点
2. 就是利用快指针走两步慢指针走一步的差距，当快指针走到最后的时候慢指针刚好走一半
*/
func (node *Node) GetMidNode() *Node {
	// 快指针
	fast := node
	// 慢指针
	low := node
	for fast.NextPtr.NextPtr != nil && low.NextPtr != nil {
		fast = fast.NextPtr.NextPtr
		low = low.NextPtr
	}
	return low
}

/*
Reverse
@brief: 链表反转
*/
func (node *Node) Reverse() *Node {
	if node == nil {
		return nil
	}
	// 前躯节点和后继节点
	pre := new(Node)
	cur := node.NextPtr
	for cur != nil {
		// 1. 保存后继节点
		next := cur.NextPtr
		// 2. 后继节点与当前节点断开
		cur.NextPtr = pre
		// 3. 见当前节点赋予断开
		pre = cur
		// 4. 往后移动
		cur = next
	}

	return pre
}
