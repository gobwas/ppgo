#include "ppgo/algorithm/sort.h"

/**
 * This file contains an implementation of immutable sorted array with go's
 * array as a backend.
 */

#ifndef _PPGO_STRUCT_TUPLE_
#define _PPGO_STRUCT_TUPLE_

#define MAKE_TUPLE(N, T, K);;\
const VAR(Capacity) = N;;\
;;\
type STRUCT() struct {;;\
	data ARRAY(VAR(Capacity), T);;\
	size int;;\
};;\
;;\
func (t STRUCT()) Has(x K) bool {;;\
	DO_SEARCH_RANGE(t.data, x, 0, t.size, i, ok);;\
	return ok;;\
};;\
;;\
func (t STRUCT()) Get(x K) (T, bool) {;;\
	DO_SEARCH_RANGE(t.data, x, 0, t.size, i, ok);;\
	if !ok {;;\
		return EMPTY(), false;;\
	};;\
	return t.data[i], true;;\
};;\
>>> Upsert inserts item x into array or updates existing one.;;\
>>> It returns copy of STRUCT(), previous item (if were present) and a boolean;;\
>>> flag that reports about previous item replacement. This flag is useful for;;\
>>> non-pointer item types such as numbers or struct values.;;\
>>> ;;\
>>> Note that it will panic on out of range insertion.;;\
func (t STRUCT()) Upsert(x T) (cp STRUCT(), prev T, ok bool) {;;\
	DO_SEARCH_RANGE(t.data, ID(x), 0, t.size, i, has);;\
	if has {;;\
		t.data[i], prev = x, t.data[i];;\
		ok = true;;\
	} else {;;\
		t.size++;;\
		copy(t.data[i+1:t.size], t.data[i:t.size-1]);;\
		t.data[i] = x;;\
		prev = EMPTY();;\
	};;\
	return t, prev, ok;;\
};;\
;;\
func (t STRUCT()) Delete(x K) (cp STRUCT(), prev T, ok bool) {;;\
	DO_SEARCH_RANGE(t.data, ID(x), 0, t.size, i, has);;\
	if !has {;;\
		return t, EMPTY(), false;\
	};;\
	t.size--;;\
	prev = t.data[i];;\
	copy(t.data[i:t.size], t.data[i+1:t.size+1]);;\
	return t, prev, true;;\
};;\
;;\
func (t STRUCT()) Ascend(cb func(x T) bool) bool {;;\
	for i := 0; i < t.size; i++ {;;\
		if !cb(t.data[i]) {;;\
			return false;;\
		};;\
	};;\
	return true;;\
};;\
;;\
func (t STRUCT()) AscendRange(x, y K, cb func(x T) bool) bool {;;\
	DO_SEARCH_RANGE(t.data, x, 0, t.size, i, hasX);;\
	DO_SEARCH_RANGE(t.data, y, i, t.size, j, hasY);;\
	for ; i < t.size && i <= j; i++ {;;\
		if !cb(t.data[i]) {;;\
			return false;;\
		};;\
	};;\
	return true;;\
};;\
;;\
func (t STRUCT()) Reset() STRUCT() {;;\
	for i := 0; i < t.size; i++ {;;\
		>>> Need to prevent memory leaks on complex structs.;;\
		t.data[i] = EMPTY();;\
	};;\
	t.size = 0;;\
	return t;;\
};;\
;;\
func (t STRUCT()) Append(to SLICE(T)) SLICE(T) {;;\
	return append(to, t.data[:t.size]...);;\
};;\
;;\
func (t STRUCT()) Len() int {;;\
	return t.size;;\
};;\
;;\
func (t STRUCT()) Cap() int {;;\
	return VAR(Capacity) - t.size;;\
};;\
;;\

#endif /* !_PPGO_STRUCT_TUPLE_ */
