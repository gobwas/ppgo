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

func NewSyncArray() *SyncArray {
	return &SyncArray{}
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
	case r == 0: // no readers, insert inplace
		if cap(a.data) == len(a.data) { // not enough storage in array
			goto copyCase
		}
		a.data = a.data[:len(a.data)+1]
		copy(a.data[i+1:], a.data[i:])
		a.data[i] = x
	copyCase:
		fallthrough
	case r > 0: // readers exists, do copy
		with := make([]int, len(a.data)+1)
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
	case r == 0: // no readers, insert inplace
		if cap(a.data) == len(a.data) { // not enough storage in array
			goto copyCase
		}
		a.data = a.data[:len(a.data)+1]
		copy(a.data[i+1:], a.data[i:])
		a.data[i] = x
	copyCase:
		fallthrough
	case r > 0: // readers exists, do copy
		with := make([]int, len(a.data)+1)
		copy(with[:i], a.data[:i])
		copy(with[i+1:], a.data[i:])
		with[i] = x
		a.data = with
	}
	a.mu.Unlock()
	return x
}
func (a *SyncArray) Upsert(x int) (prev int) {
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
	case r > 0 && has: // readers exists, do copy
		with := make([]int, len(a.data))
		copy(with, a.data)
		a.data = with
		fallthrough
	case r == 0 && has: // no readers: update in place
		a.data[i], prev = x, a.data[i]
	case r == 0 && !has: // no readers, insert inplace
		if cap(a.data) == len(a.data) { // not enough storage in array
			goto copyCase
		}
		a.data = a.data[:len(a.data)+1]
		copy(a.data[i+1:], a.data[i:])
		a.data[i] = x
	copyCase:
		fallthrough
	case r > 0 && !has: // readers exists, do copy
		with := make([]int, len(a.data)+1)
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

func (a *SyncArray) Delete(x int) (int, bool) {
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
	prev := a.data[i]
	r := atomic.LoadInt64(&a.readers)
	switch {
	case r == 0: // no readers, delete inplace
		a.data[i] = 0
		a.data = a.data[:i+copy(a.data[i:], a.data[i+1:])]
	case r > 0: // has readers, copy
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