package heapSort

type Heap []int

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) up(i int) {
	for {
		f := (i - 1) / 2
		if i == f || h.less(f, i) {
			break
		}
		h.swap(f, i)
		i = f
	}
}

func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

func (h Heap) down(i int) {
	for {
		l := 2*i + 1
		if l >= len(h) {
			break
		}
		j := l
		if r := l + 1; r < len(h) && h.less(r, l) {
			j = r
		}
		if h.less(i, j) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

func (h *Heap) Remove(i int) (int, bool) {
	if i < 0 || i > len(*h)-1 {
		return 0, false
	}
	n := len(*h) - 1
	h.swap(i, n) // 用最后的元素值替换被删除元素
	// 删除最后的元素
	x := (*h)[n]
	*h = (*h)[0:n]
	// 如果当前元素大于父结点，向下筛选
	if (*h)[i] > (*h)[(i-1)/2] {
		h.down(i)
	} else { // 当前元素小于父结点，向上筛选
		h.up(i)
	}
	return x, true
}

func (h *Heap) Pop() int {
	n := len(*h) - 1
	h.swap(0, n)
	x := (*h)[n]
	*h = (*h)[0:n]
	h.down(0)
	return x
}

func (h Heap) Init() {
	n := len(h)
	// i > n/2-1 的结点为叶子结点本身已经是堆了
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

func HeapSort(arr []int) (res []int) {
	h := Heap(arr)
	h.Init()
	for len(h) != 0 {
		res = append(res, h.Pop())
	}
	return res
}
