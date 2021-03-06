// THIS FILE WAS AUTOGENERATED.
// DO NOT EDIT!

package ints

type SortedSlice struct {
	Data []int
}

// NewSortedSlice creates SortedSlice with underlying data.
// Note that data is not copied and used by reference.
func NewSortedSlice(data []int) SortedSlice {
	_SortedSliceSortSource(data, 0, len(data))
	return SortedSlice{Data: data}
}

// _SortedSliceSortSource sorts data for further use inside SortedSlice.
func _SortedSliceSortSource(data []int, lo, hi int) {
	if hi-lo <= 12 {
		// Do insertion sort.
		for i := lo + 1; i < hi; i++ {
			for j := i; j > lo && !(data[j-1] <= data[j]); j-- {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
		return
	}
	// Do quick sort.
	var (
		p = lo
		x = data[lo]
	)
	for i := lo + 1; i < hi; i++ {
		if data[i] <= x {
			p++
			data[p], data[i] = data[i], data[p]
		}
	}
	data[p], data[lo] = data[lo], data[p]

	if lo < p {
		_SortedSliceSortSource(data, lo, p)
	}
	if p+1 < hi {
		_SortedSliceSortSource(data, p+1, hi)
	}
}

func (a SortedSlice) Has(x int) bool {
	// Binary search algorithm.
	var ok bool
	var i int
	{
		l := 0
		r := len(a.Data)
		for !ok && l < r {
			m := l + (r-l)/2
			switch {
			case a.Data[m] == x:
				ok = true
				r = m
			case a.Data[m] < x:
				l = m + 1
			case a.Data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	return ok
}

func (a SortedSlice) Get(x int) (int, bool) {
	// Binary search algorithm.
	var ok bool
	var i int
	{
		l := 0
		r := len(a.Data)
		for !ok && l < r {
			m := l + (r-l)/2
			switch {
			case a.Data[m] == x:
				ok = true
				r = m
			case a.Data[m] < x:
				l = m + 1
			case a.Data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	if !ok {
		return 0, false
	}
	return a.Data[i], true
}

// Upsert inserts item x into array or updates existing one.
// It returns copy of SortedSlice, previous item (if were present) and a boolean
// flag that reports about previous item replacement. This flag is useful for
// non-pointer item types such as numbers or struct values.
func (a SortedSlice) Upsert(x int) (cp SortedSlice, prev int, swapped bool) {
	var with []int
	// Binary search algorithm.
	var has bool
	var i int
	{
		l := 0
		r := len(a.Data)
		for !has && l < r {
			m := l + (r-l)/2
			switch {
			case a.Data[m] == x:
				has = true
				r = m
			case a.Data[m] < x:
				l = m + 1
			case a.Data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	if has {
		with = make([]int, len(a.Data))
		copy(with, a.Data)
		with[i], prev = x, a.Data[i]
		swapped = true
	} else {
		with = make([]int, len(a.Data)+1)
		copy(with[:i], a.Data[:i])
		copy(with[i+1:], a.Data[i:])
		with[i] = x
		prev = 0
	}
	return SortedSlice{with}, prev, swapped
}

func (a SortedSlice) Delete(x int) (SortedSlice, int, bool) {
	// Binary search algorithm.
	var has bool
	var i int
	{
		l := 0
		r := len(a.Data)
		for !has && l < r {
			m := l + (r-l)/2
			switch {
			case a.Data[m] == x:
				has = true
				r = m
			case a.Data[m] < x:
				l = m + 1
			case a.Data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	if !has {
		return a, 0, false
	}
	without := make([]int, len(a.Data)-1)
	copy(without[:i], a.Data[:i])
	copy(without[i:], a.Data[i+1:])
	return SortedSlice{without}, a.Data[i], true
}

func (a SortedSlice) Ascend(cb func(x int) bool) bool {
	for _, x := range a.Data {
		if !cb(x) {
			return false
		}
	}
	return true
}

func (a SortedSlice) AscendRange(x, y int, cb func(x int) bool) bool {
	// Binary search algorithm.
	var hasX bool
	var i int
	{
		l := 0
		r := len(a.Data)
		for !hasX && l < r {
			m := l + (r-l)/2
			switch {
			case a.Data[m] == x:
				hasX = true
				r = m
			case a.Data[m] < x:
				l = m + 1
			case a.Data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	// Binary search algorithm.
	var hasY bool
	var j int
	{
		l := i
		r := len(a.Data)
		for !hasY && l < r {
			m := l + (r-l)/2
			switch {
			case a.Data[m] == y:
				hasY = true
				r = m
			case a.Data[m] < y:
				l = m + 1
			case a.Data[m] > y:
				r = m
			}
		}
		j = r
		_ = j // in case when j not being used
	}
	for ; i < len(a.Data) && i <= j; i++ {
		if !cb(a.Data[i]) {
			return false
		}
	}
	return true
}

func (a SortedSlice) Reset() SortedSlice {
	return SortedSlice{nil}
}

func (a SortedSlice) AppendTo(p []int) []int {
	return append(p, a.Data...)
}
