package ints

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	for i, test := range []struct {
		before []int
		after  []int
	}{
		{
			before: []int{3, 2, 1, 0},
			after:  []int{0, 1, 2, 3},
		},
		{
			before: []int{0, -1, -2, -3},
			after:  []int{-3, -2, -1, 0},
		},
		{
			before: []int{},
			after:  []int{},
		},
	} {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			s := make([]int, len(test.before))
			copy(s, test.before)
			Sort(s, 0, len(s))
			if !reflect.DeepEqual(s, test.after) {
				t.Errorf("Sort(%v) ~> %v; want %v", test.before, s, test.after)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	for i, test := range []struct {
		data  []int
		need  int
		ok    bool
		index int
	}{
		{
			data:  []int{0, 1, 2, 3},
			need:  1,
			ok:    true,
			index: 1,
		},
		{
			data:  []int{0, 1, 3, 4},
			need:  2,
			ok:    false,
			index: 2,
		},
	} {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			index, ok := Search(test.data, test.need)
			if ok != test.ok || index != test.index {
				t.Errorf(
					"Search(%v, %v) = %v, %v; want %v, %v",
					test.data, test.need, index, ok, test.index, test.ok,
				)
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	for _, bench := range []struct {
		label string
		size  int
		fn    func([]int, int, int)
	}{
		{
			label: "ppgo",
			size:  10,
			fn:    Sort,
		},
		{
			label: "ppgo",
			size:  100,
			fn:    Sort,
		},
		{
			label: "ppgo",
			size:  1000,
			fn:    Sort,
		},
		{
			label: "golang",
			size:  10,
			fn:    func(data []int, l, r int) { sort.Ints(data) },
		},
		{
			label: "golang",
			size:  100,
			fn:    func(data []int, l, r int) { sort.Ints(data) },
		},
		{
			label: "golang",
			size:  1000,
			fn:    func(data []int, l, r int) { sort.Ints(data) },
		},
	} {
		b.Run(fmt.Sprintf("%s_%d", bench.label, bench.size), func(b *testing.B) {
			b.StopTimer()
			data := make([]int, bench.size)
			for i := 0; i < len(data); i++ {
				data[i] = i
			}
			for i := 0; i < b.N; i++ {
				for j := 0; j < len(data); j++ {
					k := rand.Intn(j + 1)
					data[j], data[k] = data[k], data[j]
				}
				b.StartTimer()
				bench.fn(data, 0, len(data))
				b.StopTimer()
			}
		})
	}
}

func BenchmarkSearch(b *testing.B) {
	for _, bench := range []struct {
		label string
		data  []int
		need  int
		fn    func([]int, int) (int, bool)
	}{
		{
			label: "ppgo",
			data:  []int{-10, -5, 0, 1, 2, 3, 5, 7, 11, 100, 100, 100, 1000, 10000},
			need:  11,
			fn:    Search,
		},
		{
			label: "golang",
			data:  []int{-10, -5, 0, 1, 2, 3, 5, 7, 11, 100, 100, 100, 1000, 10000},
			need:  11,
			fn: func(data []int, need int) (int, bool) {
				id := sort.SearchInts(data, need)
				ok := id < len(data) && data[id] == need
				return id, ok
			},
		},
	} {
		b.Run(fmt.Sprintf("%s", bench.label), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = bench.fn(bench.data, bench.need)
			}
		})
	}
}
