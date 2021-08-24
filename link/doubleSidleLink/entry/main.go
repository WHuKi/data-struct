package main

import (
	"data-struct/link/doubleSidleLink"
	"fmt"
)

func main() {
	node := doubleSidleLink.NewNode()
	node.InsertNodeOfHead(123)
	node.InsertNodeOfHead(456)
	node.InsertNodeOfHead(89)
	node.TraverNode()
	fmt.Println("........")
	node.DelNode(89)
	node.TraverNode()
}
