package ints

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIndexedHeapPush(t *testing.T) {
	for i, test := range []struct {
		d      int
		insert []int
		weight []int
		expect []recordIndexedHeap
		sorted []int
	}{
		{
			d:      2,
			insert: []int{0, 2, 1},
			weight: []int{0, 2, 1},
			sorted: []int{0, 1, 2},
			expect: []recordIndexedHeap{{0, 0}, {2, 2}, {1, 1}},
		},
		{
			d:      4,
			insert: []int{5, 0, 9, 2, 1, 3},
			weight: []int{5, 0, 9, 2, 1, 3},
			sorted: []int{0, 1, 2, 3, 5, 9},
			expect: []recordIndexedHeap{{0, 0}, {3, 3}, {9, 9}, {2, 2}, {1, 1}, {5, 5}},
		},
	} {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			h := NewIndexedHeap(test.d)
			for i, v := range test.insert {
				h.Push(v, test.weight[i])
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
