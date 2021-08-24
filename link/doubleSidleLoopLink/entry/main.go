package main

import "data-struct/link/doubleSidleLoopLink"

func main() {
	node := doubleSidleLoopLink.NewNode()
	node.InsertNodeOfHead(123).InsertNodeOfHead(345).InsertNodeOfHead(3456)
	node.TraverNode()
}
