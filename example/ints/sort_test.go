package ints

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

var sorters = []struct {
	label  string
	sorter func([]int, int, int)
}{
	{"ppgo", Sort},
	{"std", stdSort},
}

var sortFixtures = []struct {
	in  []int
	exp []int
}{
	{
		in:  []int{7, 0, 1, -1, 3},
		exp: []int{-1, 0, 1, 3, 7},
	},
	{
		in:  rand.Perm(10),
		exp: seq(10),
	},
	{
		in:  rand.Perm(50),
		exp: seq(50),
	},
	{
		in:  rand.Perm(100),
		exp: seq(100),
	},
	{
		in:  rand.Perm(1000),
		exp: seq(1000),
	},
}

func TestSort(t *testing.T) {
	for _, test := range sortFixtures {
		t.Run(fmt.Sprintf("%d", len(test.in)), func(t *testing.T) {
			act := copySet(test.in)
			Sort(act, 0, len(test.in))
			if !reflect.DeepEqual(act, test.exp) {
				t.Errorf("Sort(%v) = %v; want %v", test.in, act, test.exp)
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	for _, s := range sorters {
		for _, test := range sortFixtures {
			b.Run(fmt.Sprintf("%s(%d)", s.label, len(test.in)), func(b *testing.B) {
				data := make([][]int, 1000)
				for i := range data {
					data[i] = copySet(test.in)
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					j := i % 1000
					if j == 0 {
						b.StopTimer()
						for i := range data {
							copy(data[i], test.in)
						}
						b.StartTimer()
					}

					s.sorter(data[j], 0, len(data[j]))
				}
			})
		}
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

func copySet(data []int) []int {
	return append(([]int)(nil), data...)
}

func stdSort(data []int, lo, hi int) {
	sort.Ints(data)
}

func seq(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	return data
}
