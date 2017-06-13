// THIS FILE WAS AUTOGENERATED.
// DO NOT EDIT!

package ints

type SortedArray struct {
	data []int
}

func (a SortedArray) Has(x int) bool {
	// Binary search algorithm.
	var ok bool
	var i int
	{
		l := 0
		r := len(a.data)
		for !ok && l < r {
			m := l + (r-l)/2
			switch {
			case a.data[m] == x:
				ok = true
				r = m
			case a.data[m] < x:
				l = m + 1
			case a.data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	return ok
}

func (a SortedArray) Get(x int) (int, bool) {
	// Binary search algorithm.
	var ok bool
	var i int
	{
		l := 0
		r := len(a.data)
		for !ok && l < r {
			m := l + (r-l)/2
			switch {
			case a.data[m] == x:
				ok = true
				r = m
			case a.data[m] < x:
				l = m + 1
			case a.data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	if !ok {
		return 0, false
	}
	return a.data[i], true
}

// Upsert inserts item x into array or updates existing one.
// It returns copy of SortedArray, previous item (if were present) and a boolean
// flag that reports about previous item replacement. This flag is useful for
// non-pointer item types such as numbers or struct values.
func (a SortedArray) Upsert(x int) (cp SortedArray, prev int, ok bool) {
	var with []int
	// Binary search algorithm.
	var has bool
	var i int
	{
		l := 0
		r := len(a.data)
		for !has && l < r {
			m := l + (r-l)/2
			switch {
			case a.data[m] == x:
				has = true
				r = m
			case a.data[m] < x:
				l = m + 1
			case a.data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	if has {
		with = make([]int, len(a.data))
		copy(with, a.data)
		with[i], prev = x, a.data[i]
		ok = true
	} else {
		with = make([]int, len(a.data)+1)
		copy(with[:i], a.data[:i])
		copy(with[i+1:], a.data[i:])
		with[i] = x
		prev = 0
	}
	return SortedArray{with}, prev, ok
}

func (a SortedArray) Delete(x int) (SortedArray, int, bool) {
	// Binary search algorithm.
	var has bool
	var i int
	{
		l := 0
		r := len(a.data)
		for !has && l < r {
			m := l + (r-l)/2
			switch {
			case a.data[m] == x:
				has = true
				r = m
			case a.data[m] < x:
				l = m + 1
			case a.data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	if !has {
		return a, 0, false
	}
	without := make([]int, len(a.data)-1)
	copy(without[:i], a.data[:i])
	copy(without[i:], a.data[i+1:])
	return SortedArray{without}, a.data[i], true
}

func (a SortedArray) Ascend(cb func(x int) bool) bool {
	for _, x := range a.data {
		if !cb(x) {
			return false
		}
	}
	return true
}

func (a SortedArray) AscendRange(x, y int, cb func(x int) bool) bool {
	// Binary search algorithm.
	var hasX bool
	var i int
	{
		l := 0
		r := len(a.data)
		for !hasX && l < r {
			m := l + (r-l)/2
			switch {
			case a.data[m] == x:
				hasX = true
				r = m
			case a.data[m] < x:
				l = m + 1
			case a.data[m] > x:
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
		r := len(a.data)
		for !hasY && l < r {
			m := l + (r-l)/2
			switch {
			case a.data[m] == y:
				hasY = true
				r = m
			case a.data[m] < y:
				l = m + 1
			case a.data[m] > y:
				r = m
			}
		}
		j = r
		_ = j // in case when j not being used
	}
	for ; i < len(a.data) && i <= j; i++ {
		if !cb(a.data[i]) {
			return false
		}
	}
	return true
}

func (a SortedArray) Reset() SortedArray {
	return SortedArray{nil}
}

func (a SortedArray) Append(to []int) []int {
	return append(to, a.data...)
}

func (a SortedArray) Len() int {
	return len(a.data)
}

func (a SortedArray) Cap() int {
	return cap(a.data)
}
