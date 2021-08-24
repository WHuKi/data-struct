package main

import (
	"data-struct/link/singleLink"
	"fmt"
)

func main() {
	node := singleLink.CreateHeadNode()
	node.InsertNode(12)
	node.InsertNode(34)
	node.InsertNode(56)
	node.TraverseLink()
	fmt.Println("......")
	n := node.Reverse()
	n.TraverseLink()
}
