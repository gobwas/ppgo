// THIS FILE WAS AUTOGENERATED.
// DO NOT EDIT!

package ints

type recordIndexedHeap struct {
	x int
	w int
}

type IndexedHeap struct {
	d     int
	data  []recordIndexedHeap
	index map[int]int
}

func NewIndexedHeap(d int) *IndexedHeap {
	return &IndexedHeap{
		d:     d,
		index: make(map[int]int),
	}
}

func NewIndexedHeapFromSlice(data []int, d int) *IndexedHeap {
	records := make([]recordIndexedHeap, len(data))
	for i, x := range data {
		records[i] = recordIndexedHeap{x: x}
	}
	h := &IndexedHeap{
		d:     d,
		data:  records,
		index: make(map[int]int),
	}
	for i := len(h.data)/h.d - 1; i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}

func (h *IndexedHeap) Top() int {
	return h.data[0].x
}

func (h *IndexedHeap) Pop() int {
	n := len(h.data)
	ret := h.data[0].x
	a, b := h.data[0], h.data[n-1]
	h.index[a.x], h.index[b.x] = n-1, 0
	h.data[0], h.data[n-1] = h.data[n-1], h.data[0]
	h.data[n-1] = recordIndexedHeap{}
	h.data = h.data[:n-1]
	delete(h.index, ret)
	h.siftDown(0)
	return ret
}

func (h *IndexedHeap) Slice() []int {
	cp := *h
	cp.data = make([]recordIndexedHeap, len(h.data))
	copy(cp.data, h.data)
	cp.index = make(map[int]int) // prevent reordering original index
	ret := make([]int, len(cp.data))
	n := len(cp.data)
	for i := 0; i < n; i++ {
		ret[i] = cp.data[0].x
		cp.data = cp.data[1:]
		cp.siftDown(0)
	}
	return ret
}

func (h *IndexedHeap) Len() int { return len(h.data) }

func (h *IndexedHeap) Insert(x int, w int) {
	r := recordIndexedHeap{x, w}
	i := len(h.data)
	if cap(h.data) == len(h.data) {
		h.data = append(h.data, r)
	} else {
		h.data = h.data[:i+1]
		h.data[i] = r
	}
	h.index[x] = i
	h.siftUp(i)
}

func (h *IndexedHeap) Heapify() {
	for i := len(h.data)/h.d - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *IndexedHeap) Add(x int, w int) {
	i, ok := h.index[x]
	if !ok {
		panic("could not update value that is not present in heap")
	}
	h.update(i, recordIndexedHeap{x, h.data[i].w + w})
}

func (h *IndexedHeap) Change(x int, w int) {
	i, ok := h.index[x]
	if !ok {
		panic("could not update value that is not present in heap")
	}
	h.update(i, recordIndexedHeap{x, w})
}

func (h *IndexedHeap) Compare(a, b int) int {
	var i, j int
	i, ok := h.index[a]
	if ok {
		j, ok = h.index[b]
	}
	if !ok {
		panic("comparing record that not in heap")
	}
	return h.data[i].w - h.data[j].w
}

func (h *IndexedHeap) Remove(x int) {
	i, ok := h.index[x]
	if !ok {
		return
	}
	h.siftTop(i)
	h.Pop()
}

func (h *IndexedHeap) update(i int, r recordIndexedHeap) {
	prev := h.data[i]
	h.data[i] = r
	if !(r.w <= prev.w) {
		h.siftDown(i)
	} else {
		h.siftUp(i)
	}
}

func (h IndexedHeap) siftDown(root int) {
	for {
		min := root
		for i := 1; i <= h.d; i++ {
			child := h.d*root + i
			if child >= len(h.data) { // out of bounds
				break
			}
			if !(h.data[min].w <= h.data[child].w) {
				min = child
			}
		}
		if min == root {
			return
		}
		a, b := h.data[root], h.data[min]
		h.index[a.x], h.index[b.x] = min, root
		h.data[root], h.data[min] = h.data[min], h.data[root]
		root = min
	}
}

func (h IndexedHeap) siftUp(root int) {
	for root > 0 {
		parent := (root - 1) / h.d
		if h.data[root].w <= h.data[parent].w {
			return
		}
		a, b := h.data[parent], h.data[root]
		h.index[a.x], h.index[b.x] = root, parent
		h.data[parent], h.data[root] = h.data[root], h.data[parent]
		root = parent
	}
}

func (h IndexedHeap) siftTop(root int) {
	for root > 0 {
		parent := (root - 1) / h.d
		a, b := h.data[parent], h.data[root]
		h.index[a.x], h.index[b.x] = root, parent
		h.data[parent], h.data[root] = h.data[root], h.data[parent]
		root = parent
	}
}