package ints

import (
	"container/heap"
	"fmt"
	"reflect"
	"testing"
)

func TestHeapPush(t *testing.T) {
	for _, test := range []struct {
		d      int
		insert []int
		expect []int
		sorted []int
	}{
		{
			d:      2,
			insert: []int{0, 2, 1},
			expect: []int{0, 2, 1},
			sorted: []int{0, 1, 2},
		},
		{
			d:      4,
			insert: []int{5, 0, 9, 2, 1, 3},
			expect: []int{0, 3, 9, 2, 1, 5},
			sorted: []int{0, 1, 2, 3, 5, 9},
		},
	} {
		t.Run("", func(t *testing.T) {
			h := NewHeap(test.d)
			for _, v := range test.insert {
				h.Push(v)
			}
			if !reflect.DeepEqual(h.data, test.expect) {
				t.Errorf(
					"heap data is %v after insertion of %v into %d-ary heap; want %v",
					h.data, test.insert, test.d, test.expect,
				)
			}
			slice := h.Slice()
			if !reflect.DeepEqual(slice, test.sorted) {
				t.Errorf(
					"Slice() is %v after insertion of %v into %d-ary heap; want %v",
					slice, test.insert, test.d, test.sorted,
				)
			}
			sorted := make([]int, len(test.insert))
			for i := range test.insert {
				sorted[i] = h.Pop()
			}
			if !reflect.DeepEqual(sorted, test.sorted) {
				t.Errorf(
					"sorted is %v after insertion of %v into %d-ary heap; want %v",
					sorted, test.insert, test.d, test.sorted,
				)
			}
		})
	}
}

func TestHeapAscend(t *testing.T) {
	for _, test := range []struct {
		d      int
		insert []int
		expect []int
	}{
		{
			d:      2,
			insert: []int{0, 2, 1},
			expect: []int{0, 1, 2},
		},
		{
			d:      4,
			insert: []int{5, 0, 9, 2, 1, 3},
			expect: []int{0, 1, 2, 3, 5, 9},
		},
	} {
		t.Run("", func(t *testing.T) {
			h := NewHeap(test.d)
			for _, v := range test.insert {
				h.Push(v)
			}

			// Copy data state to check that after Ascend heap is restored.
			before := popAllAndRestore(h)

			var act []int
			h.Ascend(func(x int) bool {
				act = append(act, x)
				return true
			})
			if !reflect.DeepEqual(act, test.expect) {
				t.Errorf(
					"Ascend() on %d-ary heap called callback with %v; want %v",
					test.d, act, test.expect,
				)
			}
			if after := popAllAndRestore(h); !reflect.DeepEqual(after, before) {
				t.Errorf(
					"heap did not restored it state after Ascend(): data is %v; want %v",
					after, before,
				)
			}
		})
	}
}

func popAllAndRestore(h *Heap) (ret []int) {
	for range h.data {
		ret = append(ret, h.Pop())
	}
	for _, v := range ret {
		h.Push(v)
	}
	return ret
}

type benchHeap interface {
	Push(x int)
	Pop() int
}

func BenchmarkHeap(b *testing.B) {
	for i, bench := range []struct {
		label string
		push  []int
		pop   int
		ctor  func() benchHeap
	}{
		{
			label: "golang",
			push:  []int{0, 2, 1, 4},
			pop:   4,
			ctor:  golangHeapCtor,
		},
		{
			label: "ppgo",
			push:  []int{0, 2, 1, 4},
			pop:   4,
			ctor:  heapCtor,
		},
	} {
		b.Run(fmt.Sprintf("[%d]%s", i, bench.label), func(b *testing.B) {
			h := bench.ctor()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, x := range bench.push {
					h.Push(x)
				}
				for i := 0; i < bench.pop; i++ {
					_ = h.Pop()
				}
			}
		})
	}
}

func heapCtor() benchHeap {
	return NewHeap(2)
}

type golangHeap struct {
	h *IntHeap
}

func golangHeapCtor() benchHeap {
	h := &IntHeap{}
	heap.Init(h)
	return golangHeap{h}
}

func (g golangHeap) Push(x int) { heap.Push(g.h, x) }
func (g golangHeap) Pop() int   { return heap.Pop(g.h).(int) }

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
