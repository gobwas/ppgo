#include "ppgo/algorithm/sort.h"

/**
 * This file contains an implementation of synchronized sorted array.
 * It uses copy on write when there are some readers, otherway it makes
 * inplace mutations.
 */

#ifndef _PPGO_STRUCT_SYNC_ARRAY_
#define	_PPGO_STRUCT_SYNC_ARRAY_

#define MAKE_ARRAY(T, K);;\
import "sync";;\
import "sync/atomic";;\
;;\
;;\
>>> STRUCT() represents synchronized sorted array of T.;;\
>>> Note that in most cases you should store it somewhere by pointer.;;\
>>> This is needed because of non-pointer data inside, that used to syncrhonize usage.;;\
type STRUCT() struct {;;\
	mu sync.RWMutex;;\
	data SLICE(T);;\
	readers int64;;\
};;\
func CTOR()(n int) *STRUCT() {;;\
	return &STRUCT(){;;\
		data: make(SLICE(T), 0, n),;;\
	};;\
};;\
;;\
func (a *STRUCT()) Has(x K) bool {;;\
	READ_DATA(data);;\
	DO_SEARCH(data, x, i, ok);;\
	return ok;;\
};;\
;;\
func (a *STRUCT()) Get(x K) (T, bool) {;;\
	READ_DATA(data);;\
	DO_SEARCH(data, x, i, ok);;\
	if !ok {;;\
		return EMPTY(), false;;\
	};;\
	return data[i], true;;\
};;\
;;\
func (a *STRUCT()) GetAny(it func() (K, bool)) (T, bool) {;;\
	READ_DATA(data);;\
	for {;;\
		k, ok := it();;\
		if !ok {;;\
			break;;\
		};;\
		DO_SEARCH(data, k, i, has);;\
		if has {;;\
			return data[i], true;;\
		};;\
	};;\
	return EMPTY(), false;;\
};;\
;;\
func (a *STRUCT()) Getsert(x T) T {;;\
	a.mu.Lock();;\
	DO_SEARCH(a.data, ID(x), i, has);;\
	if has {;;\
		a.mu.Unlock();;\
		return a.data[i];;\
	};;\
	r := atomic.LoadInt64(&a.readers);;\
	switch {;;\
	case r == 0: >>> No readers, insert inplace.;;\
		INSERT_INPLACE(a.data, SLICE(T), i, x);;\
	case r > 0: >>> Readers exists, do copy.;;\
		INSERT_COPY(a.data, SLICE(T), i, x);;\
	};;\
	a.mu.Unlock();;\
	return x;;\
};;\
;;\
func (a *STRUCT()) GetsertFn(k K, factory func() T) T {;;\
	a.mu.Lock();;\
	DO_SEARCH(a.data, k, i, has);;\
	if has {;;\
		a.mu.Unlock();;\
		return a.data[i];;\
	};;\
	x := factory();;\
	r := atomic.LoadInt64(&a.readers);;\
	switch {;;\
	case r == 0: >>> No readers, insert inplace.;;\
		INSERT_INPLACE(a.data, SLICE(T), i, x);;\
	case r > 0: >>> Readers exists, do copy.;;\
		INSERT_COPY(a.data, SLICE(T), i, x);;\
	};;\
	a.mu.Unlock();;\
	return x;;\
};;\
;;\
func (a *STRUCT()) GetsertAnyFn(it func() (K, bool), factory func() T) T {;;\
	a.mu.Lock();;\
	for {;;\
		k, ok := it();;\
		if !ok {;;\
			break;;\
		};;\
		DO_SEARCH(a.data, k, i, has);;\
		if has {;;\
			a.mu.Unlock();;\
			return a.data[i];;\
		};;\
	};;\
	x := factory();;\
	DO_SEARCH(a.data, ID(x), i, has);;\
	if has {;;\
		panic("inserting item that is already exists");;\
	};;\
	r := atomic.LoadInt64(&a.readers);;\
	switch {;;\
	case r == 0: >>> No readers, insert inplace;;\
		INSERT_INPLACE(a.data, SLICE(T), i, x);;\
	case r > 0: >>> readers exists, do copy;;\
		INSERT_COPY(a.data, SLICE(T), i, x);;\
	};;\
	a.mu.Unlock();;\
	return x;;\
};;\
;;\
>>> Upsert inserts item x into array or updates existing one.;;\
>>> It returns previous item (if were present) and a boolean flag that reports;;\
>>> about previous item replacement. This flag is useful for non-pointer item types;;\
>>> such as numbers or struct values.;;\
func (a *STRUCT()) Upsert(x T) (prev T, ok bool) {;;\
	a.mu.Lock();;\
	DO_SEARCH(a.data, ID(x), i, has);;\
	r := atomic.LoadInt64(&a.readers);;\
	switch {;;\
	case r > 0 && has: >>> Readers exists, do copy.;;\
		with := make(SLICE(T), len(a.data));;\
		copy(with, a.data);;\
		a.data = with;;\
		fallthrough;;\
	case r == 0 && has: >>> No readers: update in place.;;\
		a.data[i], prev = x, a.data[i];;\
		ok = true;;\
	case r == 0 && !has: >>> No readers, insert inplace;;\
		INSERT_INPLACE(a.data, SLICE(T), i, x);;\
	case r > 0 && !has: >>> Readers exists, do copy.;;\
		INSERT_COPY(a.data, SLICE(T), i, x);;\
	};;\
	a.mu.Unlock();;\
	return;;\
};;\
;;\
func (a *STRUCT()) Do(cb func(SLICE(T))) {;;\
	READ_DATA(data);;\
	cb(data);;\
};;\
;;\
func (a *STRUCT()) Delete(x K) (T, bool) {;;\
	return a.DeleteCond(x, nil);;\
};;\
;;\
func (a *STRUCT()) DeleteCond(x K, predicate func(T) bool) (T, bool) {;;\
	a.mu.Lock();;\
	DO_SEARCH(a.data, x, i, has);;\
	if !has {;;\
		a.mu.Unlock();;\
		return EMPTY(), false;;\
	};;\
	if predicate != nil && !predicate(a.data[i]) {;;\
		a.mu.Unlock();;\
		return EMPTY(), false;;\
	};;\
	prev := a.data[i];;\
	r := atomic.LoadInt64(&a.readers);;\
	switch {;;\
	case r == 0: >>> No readers, delete inplace.;;\
		a.data[i] = EMPTY();;\
		a.data = a.data[:i+copy(a.data[i:], a.data[i+1:])];;\
	case r > 0: >>> Has readers, copy.;;\
		without := make(SLICE(T), len(a.data)-1);;\
		copy(without[:i], a.data[:i]);;\
		copy(without[i:], a.data[i+1:]);;\
		a.data = without;;\
	};;\
	a.mu.Unlock();;\
	return prev, true;;\
};;\
;;\
func (a *STRUCT()) Ascend(cb func(x T) bool) bool {;;\
	READ_DATA(data);;\
	for _, x := range data {;;\
		if !cb(x) {;;\
			return false;;\
		};;\
	};;\
	return true;;\
};;\
;;\
func (a *STRUCT()) AscendRange(x, y K, cb func(x T) bool) bool {;;\
	READ_DATA(data);;\
	DO_SEARCH_RANGE(a.data, x, 0, len(a.data), i, hasX);;\
	DO_SEARCH_RANGE(a.data, y, i, len(a.data), j, hasY);;\
	for ; i < len(data) && i <= j; i++ {;;\
		if !cb(data[i]) {;;\
			return false;;\
		};;\
	};;\
	return true;;\
};;\
;;\
func (a *STRUCT()) Len() int {;;\
	a.mu.RLock();;\
	n := len(a.data);;\
	a.mu.RUnlock();;\
	return n;;\
};;\

#define INSERT_INPLACE(DATA, CONTAINER, I, X)\
	if n := len(DATA); n == cap(DATA) { ;;\
		>>> No space for insertion. Grow.;;\
		DO_INSERT_COPY(DATA, CONTAINER, n*3/2 + 1, I, X);;\
	} else {;;\
		DATA = DATA[:len(DATA)+1];;\
		copy(DATA[I+1:], DATA[I:]);;\
		DATA[I] = X;;\
	}\

#define INSERT_COPY(DATA, CONTAINER, I, X)\
	grow := len(DATA) + 1;;\
	if n := len(DATA); n == cap(DATA) {;;\
		>>> No space for insertion. Grow.;;\
		grow = len(DATA)*3/2 + 1;;\
	};;\
	DO_INSERT_COPY(DATA, CONTAINER, grow, I, X)\

#define DO_INSERT_COPY(DATA, CONTAINER, GROW, I, X)\
	with := make(CONTAINER, len(DATA)+1, GROW);;\
	copy(with[:I], DATA[:I]);;\
	copy(with[I+1:], DATA[I:]);;\
	with[I] = X;;\
	DATA = with\


#define READ_DATA(DATA)\
	a.mu.RLock();;\
	DATA := a.data;;\
	atomic.AddInt64(&a.readers, 1);;\
	defer atomic.AddInt64(&a.readers, -1);;\
	a.mu.RUnlock()\

#endif /* !_PPGO_STRUCT_SYNC_ARRAY_ */
