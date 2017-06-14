#include "ppgo/algorithm/sort.h"

/**
 * This file contains an implementation of immutable sorted array.
 */

#ifndef _PPGO_STRUCT_SORTED_ARRAY_
#define _PPGO_STRUCT_SORTED_ARRAY_

#define MAKE_SORTED_ARRAY(N, T, K);;\
const VAR(Capacity) = N;;\
;;\
type STRUCT() struct {;;\
	data ARRAY(VAR(Capacity), T);;\
	size int;;\
};;\
;;\
func (a STRUCT()) Has(x K) bool {;;\
	DO_SEARCH_RANGE(a.data, x, 0, a.size, i, ok);;\
	return ok;;\
};;\
;;\
func (a STRUCT()) Get(x K) (T, bool) {;;\
	DO_SEARCH_RANGE(a.data, x, 0, a.size, i, ok);;\
	if !ok {;;\
		return EMPTY(), false;;\
	};;\
	return a.data[i], true;;\
};;\
>>> Upsert inserts item x into array or updates existing one.;;\
>>> It returns copy of STRUCT(), previous item (if were present) and a boolean;;\
>>> flag that reports about previous item replacement. This flag is useful for;;\
>>> non-pointer item types such as numbers or struct values.;;\
>>> ;;\
>>> Note that it will panic on out of range insertion.;;\
func (a STRUCT()) Upsert(x T) (cp STRUCT(), prev T, ok bool) {;;\
	DO_SEARCH_RANGE(a.data, ID(x), 0, a.size, i, has);;\
	if has {;;\
		a.data[i], prev = x, a.data[i];;\
		ok = true;;\
	} else {;;\
		a.size++;;\
		copy(a.data[i+1:a.size], a.data[i:a.size-1]);;\
		a.data[i] = x;;\
		prev = EMPTY();;\
	};;\
	return t, prev, ok;;\
};;\
;;\
func (a STRUCT()) Delete(x K) (cp STRUCT(), prev T, ok bool) {;;\
	DO_SEARCH_RANGE(a.data, ID(x), 0, a.size, i, has);;\
	if !has {;;\
		return t, EMPTY(), false;\
	};;\
	a.size--;;\
	prev = a.data[i];;\
	copy(a.data[i:a.size], a.data[i+1:a.size+1]);;\
	return t, prev, true;;\
};;\
;;\
func (a STRUCT()) Ascend(cb func(x T) bool) bool {;;\
	for i := 0; i < a.size; i++ {;;\
		if !cb(a.data[i]) {;;\
			return false;;\
		};;\
	};;\
	return true;;\
};;\
;;\
func (a STRUCT()) AscendRange(x, y K, cb func(x T) bool) bool {;;\
	DO_SEARCH_RANGE(a.data, x, 0, a.size, i, hasX);;\
	DO_SEARCH_RANGE(a.data, y, i, a.size, j, hasY);;\
	for ; i < a.size && i <= j; i++ {;;\
		if !cb(a.data[i]) {;;\
			return false;;\
		};;\
	};;\
	return true;;\
};;\
;;\
func (a STRUCT()) Reset() STRUCT() {;;\
	for i := 0; i < a.size; i++ {;;\
		>>> Need to prevent memory leaks on complex structs.;;\
		a.data[i] = EMPTY();;\
	};;\
	a.size = 0;;\
	return t;;\
};;\
;;\
func (a STRUCT()) AppendTo(p SLICE(T)) SLICE(T) {;;\
	return append(p, a.data[:a.size]...);;\
};;\
;;\
func (a STRUCT()) Len() int {;;\
	return a.size;;\
};;\
;;\
func (a STRUCT()) Cap() int {;;\
	return VAR(Capacity) - a.size;;\
};;\
;;\

#endif /* !_PPGO_STRUCT_SORTED_ARRAY_ */
