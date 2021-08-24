package main

import (
	"data-struct/link/singleLink"
	"fmt"
)

func main() {
	node := singleLink.CreateHeadNode()
	node.InsertNode(12)

	node.InsertNodeOfHead(34)
	node.InsertNode(56)
	node.TraverseLink()
	fmt.Println(".......")
	node.DelLinkNode(34)
	node.TraverseLink()
}
