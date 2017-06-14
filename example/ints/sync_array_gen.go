// THIS FILE WAS AUTOGENERATED.
// DO NOT EDIT!

package ints

import "sync"
import "sync/atomic"

// SyncArray represents synchronized sorted array of int.
// Note that in most cases you should store it somewhere by pointer.
// This is needed because of non-pointer data inside, that used to syncrhonize usage.
type SyncArray struct {
	mu      sync.RWMutex
	data    []int
	readers int64
}

func NewSyncArray(n int) *SyncArray {
	return &SyncArray{
		data: make([]int, 0, n),
	}
}

// NewSyncArrayFromSlice creates SyncArray with src as underlying data.
// Note that src is not copied and used by reference.
func NewSyncArrayFromSlice(data []int) *SyncArray {
	sortSyncArraySource(data, 0, len(data))
	return &SyncArray{
		data: data,
	}
}

// sortSyncArraySource sorts data for further use inside SyncArray.
func sortSyncArraySource(data []int, lo, hi int) {
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
		sortSyncArraySource(data, lo, p)
	}
	if p+1 < hi {
		sortSyncArraySource(data, p+1, hi)
	}
}

func (a *SyncArray) Has(x int) bool {
	a.mu.RLock()
	data := a.data
	atomic.AddInt64(&a.readers, 1)
	defer atomic.AddInt64(&a.readers, -1)
	a.mu.RUnlock()
	// Binary search algorithm.
	var ok bool
	var i int
	{
		l := 0
		r := len(data)
		for !ok && l < r {
			m := l + (r-l)/2
			switch {
			case data[m] == x:
				ok = true
				r = m
			case data[m] < x:
				l = m + 1
			case data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	return ok
}

func (a *SyncArray) Get(x int) (int, bool) {
	a.mu.RLock()
	data := a.data
	atomic.AddInt64(&a.readers, 1)
	defer atomic.AddInt64(&a.readers, -1)
	a.mu.RUnlock()
	// Binary search algorithm.
	var ok bool
	var i int
	{
		l := 0
		r := len(data)
		for !ok && l < r {
			m := l + (r-l)/2
			switch {
			case data[m] == x:
				ok = true
				r = m
			case data[m] < x:
				l = m + 1
			case data[m] > x:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	if !ok {
		return 0, false
	}
	return data[i], true
}

func (a *SyncArray) GetAny(it func() (int, bool)) (int, bool) {
	a.mu.RLock()
	data := a.data
	atomic.AddInt64(&a.readers, 1)
	defer atomic.AddInt64(&a.readers, -1)
	a.mu.RUnlock()
	for {
		k, ok := it()
		if !ok {
			break
		}
		// Binary search algorithm.
		var has bool
		var i int
		{
			l := 0
			r := len(data)
			for !has && l < r {
				m := l + (r-l)/2
				switch {
				case data[m] == k:
					has = true
					r = m
				case data[m] < k:
					l = m + 1
				case data[m] > k:
					r = m
				}
			}
			i = r
			_ = i // in case when i not being used
		}
		if has {
			return data[i], true
		}
	}
	return 0, false
}

func (a *SyncArray) Getsert(x int) int {
	a.mu.Lock()
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
		a.mu.Unlock()
		return a.data[i]
	}
	r := atomic.LoadInt64(&a.readers)
	switch {
	case r == 0: // No readers, insert inplace.
		if n := len(a.data); n == cap(a.data) {
			// No space for insertion. Grow.
			with := make([]int, len(a.data)+1, n*3/2+1)
			copy(with[:i], a.data[:i])
			copy(with[i+1:], a.data[i:])
			with[i] = x
			a.data = with
		} else {
			a.data = a.data[:len(a.data)+1]
			copy(a.data[i+1:], a.data[i:])
			a.data[i] = x
		}
	case r > 0: // Readers exists, do copy.
		grow := len(a.data) + 1
		if n := len(a.data); n == cap(a.data) {
			// No space for insertion. Grow.
			grow = len(a.data)*3/2 + 1
		}
		with := make([]int, len(a.data)+1, grow)
		copy(with[:i], a.data[:i])
		copy(with[i+1:], a.data[i:])
		with[i] = x
		a.data = with
	}
	a.mu.Unlock()
	return x
}

func (a *SyncArray) GetsertFn(k int, factory func() int) int {
	a.mu.Lock()
	// Binary search algorithm.
	var has bool
	var i int
	{
		l := 0
		r := len(a.data)
		for !has && l < r {
			m := l + (r-l)/2
			switch {
			case a.data[m] == k:
				has = true
				r = m
			case a.data[m] < k:
				l = m + 1
			case a.data[m] > k:
				r = m
			}
		}
		i = r
		_ = i // in case when i not being used
	}
	if has {
		a.mu.Unlock()
		return a.data[i]
	}
	x := factory()
	r := atomic.LoadInt64(&a.readers)
	switch {
	case r == 0: // No readers, insert inplace.
		if n := len(a.data); n == cap(a.data) {
			// No space for insertion. Grow.
			with := make([]int, len(a.data)+1, n*3/2+1)
			copy(with[:i], a.data[:i])
			copy(with[i+1:], a.data[i:])
			with[i] = x
			a.data = with
		} else {
			a.data = a.data[:len(a.data)+1]
			copy(a.data[i+1:], a.data[i:])
			a.data[i] = x
		}
	case r > 0: // Readers exists, do copy.
		grow := len(a.data) + 1
		if n := len(a.data); n == cap(a.data) {
			// No space for insertion. Grow.
			grow = len(a.data)*3/2 + 1
		}
		with := make([]int, len(a.data)+1, grow)
		copy(with[:i], a.data[:i])
		copy(with[i+1:], a.data[i:])
		with[i] = x
		a.data = with
	}
	a.mu.Unlock()
	return x
}

func (a *SyncArray) GetsertAnyFn(it func() (int, bool), factory func() int) int {
	a.mu.Lock()
	for {
		k, ok := it()
		if !ok {
			break
		}
		// Binary search algorithm.
		var has bool
		var i int
		{
			l := 0
			r := len(a.data)
			for !has && l < r {
				m := l + (r-l)/2
				switch {
				case a.data[m] == k:
					has = true
					r = m
				case a.data[m] < k:
					l = m + 1
				case a.data[m] > k:
					r = m
				}
			}
			i = r
			_ = i // in case when i not being used
		}
		if has {
			a.mu.Unlock()
			return a.data[i]
		}
	}
	x := factory()
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
		panic("inserting item that is already exists")
	}
	r := atomic.LoadInt64(&a.readers)
	switch {
	case r == 0: // No readers, insert inplace
		if n := len(a.data); n == cap(a.data) {
			// No space for insertion. Grow.
			with := make([]int, len(a.data)+1, n*3/2+1)
			copy(with[:i], a.data[:i])
			copy(with[i+1:], a.data[i:])
			with[i] = x
			a.data = with
		} else {
			a.data = a.data[:len(a.data)+1]
			copy(a.data[i+1:], a.data[i:])
			a.data[i] = x
		}
	case r > 0: // readers exists, do copy
		grow := len(a.data) + 1
		if n := len(a.data); n == cap(a.data) {
			// No space for insertion. Grow.
			grow = len(a.data)*3/2 + 1
		}
		with := make([]int, len(a.data)+1, grow)
		copy(with[:i], a.data[:i])
		copy(with[i+1:], a.data[i:])
		with[i] = x
		a.data = with
	}
	a.mu.Unlock()
	return x
}

// Upsert inserts item x into array or updates existing one.
// It returns previous item (if were present) and a boolean flag that reports
// about previous item replacement. This flag is useful for non-pointer item types
// such as numbers or struct values.
func (a *SyncArray) Upsert(x int) (prev int, ok bool) {
	a.mu.Lock()
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
	r := atomic.LoadInt64(&a.readers)
	switch {
	case r > 0 && has: // Readers exists, do copy.
		with := make([]int, len(a.data))
		copy(with, a.data)
		a.data = with
		fallthrough
	case r == 0 && has: // No readers: update in place.
		a.data[i], prev = x, a.data[i]
		ok = true
	case r == 0 && !has: // No readers, insert inplace
		if n := len(a.data); n == cap(a.data) {
			// No space for insertion. Grow.
			with := make([]int, len(a.data)+1, n*3/2+1)
			copy(with[:i], a.data[:i])
			copy(with[i+1:], a.data[i:])
			with[i] = x
			a.data = with
		} else {
			a.data = a.data[:len(a.data)+1]
			copy(a.data[i+1:], a.data[i:])
			a.data[i] = x
		}
	case r > 0 && !has: // Readers exists, do copy.
		grow := len(a.data) + 1
		if n := len(a.data); n == cap(a.data) {
			// No space for insertion. Grow.
			grow = len(a.data)*3/2 + 1
		}
		with := make([]int, len(a.data)+1, grow)
		copy(with[:i], a.data[:i])
		copy(with[i+1:], a.data[i:])
		with[i] = x
		a.data = with
	}
	a.mu.Unlock()
	return
}

func (a *SyncArray) Do(cb func([]int)) {
	a.mu.RLock()
	data := a.data
	atomic.AddInt64(&a.readers, 1)
	defer atomic.AddInt64(&a.readers, -1)
	a.mu.RUnlock()
	cb(data)
}

func (a *SyncArray) AppendTo(p []int) []int {
	a.mu.RLock()
	data := a.data
	atomic.AddInt64(&a.readers, 1)
	defer atomic.AddInt64(&a.readers, -1)
	a.mu.RUnlock()
	return append(p, data...)
}

func (a *SyncArray) Delete(x int) (int, bool) {
	return a.DeleteCond(x, nil)
}

func (a *SyncArray) DeleteCond(x int, predicate func(int) bool) (int, bool) {
	a.mu.Lock()
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
		a.mu.Unlock()
		return 0, false
	}
	if predicate != nil && !predicate(a.data[i]) {
		a.mu.Unlock()
		return 0, false
	}
	prev := a.data[i]
	r := atomic.LoadInt64(&a.readers)
	switch {
	case r == 0: // No readers, delete inplace.
		a.data[i] = 0
		a.data = a.data[:i+copy(a.data[i:], a.data[i+1:])]
	case r > 0: // Has readers, copy.
		without := make([]int, len(a.data)-1)
		copy(without[:i], a.data[:i])
		copy(without[i:], a.data[i+1:])
		a.data = without
	}
	a.mu.Unlock()
	return prev, true
}

func (a *SyncArray) Ascend(cb func(x int) bool) bool {
	a.mu.RLock()
	data := a.data
	atomic.AddInt64(&a.readers, 1)
	defer atomic.AddInt64(&a.readers, -1)
	a.mu.RUnlock()
	for _, x := range data {
		if !cb(x) {
			return false
		}
	}
	return true
}

func (a *SyncArray) AscendRange(x, y int, cb func(x int) bool) bool {
	a.mu.RLock()
	data := a.data
	atomic.AddInt64(&a.readers, 1)
	defer atomic.AddInt64(&a.readers, -1)
	a.mu.RUnlock()
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
	for ; i < len(data) && i <= j; i++ {
		if !cb(data[i]) {
			return false
		}
	}
	return true
}

func (a *SyncArray) Len() int {
	a.mu.RLock()
	n := len(a.data)
	a.mu.RUnlock()
	return n
}
