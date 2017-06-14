// THIS FILE WAS AUTOGENERATED.
// DO NOT EDIT!

package ints

const ArrayCapacity = 8

type Array struct {
	data [ArrayCapacity]int
	size int
}

// NewArray creates Array with underlying sorted copy of given data.
func NewArray(data []int) Array {
	a := Array{}
	a.size = copy(a.data, data)
	a.sort(0, a.size)
	return a
}

// sort sorts data for further use. It is intended to be
// used only once in NewArray.
func (a *Array) sort(lo, hi int) {
	if hi-lo <= 12 {
		// Do insertion sort.
		for i := lo + 1; i < hi; i++ {
			for j := i; j > lo && !(a.data[j-1] <= a.data[j]); j-- {
				a.data[j], a.data[j-1] = a.data[j-1], a.data[j]
			}
		}
		return
	}
	// Do quick sort.
	var (
		p = lo
		x = a.data[lo]
	)
	for i := lo + 1; i < hi; i++ {
		if a.data[i] <= x {
			p++
			a.data[p], a.data[i] = a.data[i], a.data[p]
		}
	}
	a.data[p], a.data[lo] = a.data[lo], a.data[p]

	if lo < p {
		a.sort(data, lo, p)
	}
	if p+1 < hi {
		a.sort(data, p+1, hi)
	}
}

func (a *Array) Has(x int) bool {
	// Binary search algorithm.
	var ok bool
	var i int
	{
		l := 0
		r := a.size
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

func (a *Array) Get(x int) (int, bool) {
	// Binary search algorithm.
	var ok bool
	var i int
	{
		l := 0
		r := a.size
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
// It returns copy of Array, previous item (if were present) and a boolean
// flag that reports about previous item replacement. This flag is useful for
// non-pointer item types such as numbers or struct values.
//
// Note that it will panic on out of range insertion.
func (a Array) Upsert(x int) (cp Array, prev int, ok bool) {
	// Binary search algorithm.
	var has bool
	var i int
	{
		l := 0
		r := a.size
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
		a.data[i], prev = x, a.data[i]
		ok = true
	} else {
		a.size++
		copy(a.data[i+1:a.size], a.data[i:a.size-1])
		a.data[i] = x
		prev = 0
	}
	return a, prev, ok
}

func (a Array) Delete(x int) (cp Array, prev int, ok bool) {
	// Binary search algorithm.
	var has bool
	var i int
	{
		l := 0
		r := a.size
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
	a.size--
	prev = a.data[i]
	copy(a.data[i:a.size], a.data[i+1:a.size+1])
	return a, prev, true
}

func (a *Array) Ascend(cb func(x int) bool) bool {
	for i := 0; i < a.size; i++ {
		if !cb(a.data[i]) {
			return false
		}
	}
	return true
}

func (a *Array) AscendRange(x, y int, cb func(x int) bool) bool {
	// Binary search algorithm.
	var hasX bool
	var i int
	{
		l := 0
		r := a.size
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
		r := a.size
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
	for ; i < a.size && i <= j; i++ {
		if !cb(a.data[i]) {
			return false
		}
	}
	return true
}

func (a Array) Reset() Array {
	for i := 0; i < a.size; i++ {
		// Need to prevent memory leaks on complex structs.
		a.data[i] = 0
	}
	a.size = 0
	return a
}

func (a *Array) AppendTo(p []int) []int {
	return append(p, a.data[:a.size]...)
}

func (a *Array) Len() int {
	return a.size
}

func (a *Array) Cap() int {
	return ArrayCapacity - a.size
}
