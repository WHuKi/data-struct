package binary_search_tree

type binarySearchTree struct {
	value       int
	left, right *binarySearchTree
}

//NewBinarySearchTree 初始化
func NewBinarySearchTree(value int) *binarySearchTree {
	return &binarySearchTree{
		value: value,
	}
}

//Insert 添加元素
func (b *binarySearchTree) Insert(value int) *binarySearchTree {
	if b == nil {
		b = NewBinarySearchTree(value)
		return b
	}

	if b.value > value {
		b.left = b.Insert(value)
	} else {
		b.right = b.Insert(value)
	}

	return b
}

//Contain 是否包含某个元素
func (b *binarySearchTree) Contain(value int) bool {
	if b == nil {
		return false
	}

	v := b.compareTo(value)
	if v < 0 {
		b.right.Contain(value)
	} else if v > 0 {
		b.left.Contain(value)
	} else {
		return true
	}

	return false
}

func (b *binarySearchTree) compareTo(value int) int {
	return b.value - value
}

//findMax 查找最大值
func (b *binarySearchTree) findMax() int {
	if b == nil {
		return -1
	}

	if b.right == nil {
		return b.value
	} else {
		return b.right.findMax()
	}
}

//findMaxNode 查找最大节点
func (b *binarySearchTree) findMaxNode() *binarySearchTree {
	if b != nil {
		for b.right != nil {
			b = b.right
		}
	}

	return b
}
