// THIS FILE WAS AUTOGENERATED.
// DO NOT EDIT!

package ints

const TupleCapacity = 8

type Tuple struct {
	data [8]int
	size int
}

func (t Tuple) Has(x int) bool {
	// Binary search algorithm.
	var ok bool
	var i int
	{
		l := 0
		r := t.size
		for !ok && l < r {
			m := l + (r-l)/2
			switch {
			case t.data[m] == x:
				ok = true
				r = m
			case t.data[m] < x:
				l = m + 1
			case t.data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	return ok
}

func (t Tuple) Get(x int) (int, bool) {
	// Binary search algorithm.
	var ok bool
	var i int
	{
		l := 0
		r := t.size
		for !ok && l < r {
			m := l + (r-l)/2
			switch {
			case t.data[m] == x:
				ok = true
				r = m
			case t.data[m] < x:
				l = m + 1
			case t.data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	if !ok {
		return 0, false
	}
	return t.data[i], true
}

// Upsert inserts item x into array or updates existing one.
// It returns copy of Tuple, previous item (if were present) and a boolean
// flag that reports about previous item replacement. This flag is useful for
// non-pointer item types such as numbers or struct values.
//
// Note that it will panic on out of range insertion.
func (t Tuple) Upsert(x int) (cp Tuple, prev int, ok bool) {
	// Binary search algorithm.
	var has bool
	var i int
	{
		l := 0
		r := t.size
		for !has && l < r {
			m := l + (r-l)/2
			switch {
			case t.data[m] == x:
				has = true
				r = m
			case t.data[m] < x:
				l = m + 1
			case t.data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	if has {
		t.data[i], prev = x, t.data[i]
		ok = true
	} else {
		t.size++
		copy(t.data[i+1:t.size], t.data[i:t.size-1])
		t.data[i] = x
		prev = 0
	}
	return t, prev, ok
}

func (t Tuple) Delete(x int) (cp Tuple, prev int, ok bool) {
	// Binary search algorithm.
	var has bool
	var i int
	{
		l := 0
		r := t.size
		for !has && l < r {
			m := l + (r-l)/2
			switch {
			case t.data[m] == x:
				has = true
				r = m
			case t.data[m] < x:
				l = m + 1
			case t.data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	if !has {
		return t, 0, false
	}
	t.size--
	prev = t.data[i]
	copy(t.data[i:t.size], t.data[i+1:t.size+1])
	return t, prev, true
}

func (t Tuple) Ascend(cb func(x int) bool) bool {
	for i := 0; i < t.size; i++ {
		if !cb(t.data[i]) {
			return false
		}
	}
	return true
}

func (t Tuple) AscendRange(x, y int, cb func(x int) bool) bool {
	// Binary search algorithm.
	var hasX bool
	var i int
	{
		l := 0
		r := t.size
		for !hasX && l < r {
			m := l + (r-l)/2
			switch {
			case t.data[m] == x:
				hasX = true
				r = m
			case t.data[m] < x:
				l = m + 1
			case t.data[m] > x:
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
		r := t.size
		for !hasY && l < r {
			m := l + (r-l)/2
			switch {
			case t.data[m] == y:
				hasY = true
				r = m
			case t.data[m] < y:
				l = m + 1
			case t.data[m] > y:
				r = m
			}
		}
		j = r
		_ = j // in case when j not being used
	}
	for ; i < t.size && i <= j; i++ {
		if !cb(t.data[i]) {
			return false
		}
	}
	return true
}

func (t Tuple) Reset() Tuple {
	for i := 0; i < t.size; i++ {
		// Need to prevent memory leaks on complex structs.
		t.data[i] = 0
	}
	t.size = 0
	return t
}

func (t Tuple) Append(to []int) []int {
	return append(to, t.data[:t.size]...)
}

func (t Tuple) Len() int {
	return t.size
}

func (t Tuple) Cap() int {
	return 8 - t.size
}