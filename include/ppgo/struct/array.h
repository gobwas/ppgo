#include "ppgo/algorithm/sort.h"

/**
 * This file contains an implementation of immutable sorted array with go's
 * slice as a backend.
 */

#ifndef _PPGO_STRUCT_ARRAY_
#define	_PPGO_STRUCT_ARRAY_

#define MAKE_ARRAY(T, K);;\
type STRUCT() struct {;;\
	data SLICE(T);;\
};;\
;;\
func (a STRUCT()) Has(x K) bool {;;\
	DO_SEARCH(a.data, ID(x), i, ok);;\
	return ok;;\
};;\
;;\
func (a STRUCT()) Get(x K) (T, bool) {;;\
	DO_SEARCH(a.data, ID(x), i, ok);;\
	if !ok {;;\
		return EMPTY(), false;;\
	};;\
	return a.data[i], true;;\
};;\
;;\
>>> Upsert inserts item x into array or updates existing one.;;\
>>> It returns copy of STRUCT(), previous item (if were present) and a boolean;;\
>>> flag that reports about previous item replacement. This flag is useful for;;\
>>> non-pointer item types such as numbers or struct values.;;\
func (a STRUCT()) Upsert(x T) (cp STRUCT(), prev T, ok bool) {;;\
	var with SLICE(T);;\
	DO_SEARCH(a.data, ID(x), i, has);;\
	if has {;;\
		with = make(SLICE(T), len(a.data));;\
		copy(with, a.data);;\
		with[i], prev = x, a.data[i];;\
		ok = true;;\
	} else {;;\
		with = make(SLICE(T), len(a.data)+1);;\
		copy(with[:i], a.data[:i]);;\
		copy(with[i+1:], a.data[i:]);;\
		with[i] = x;;\
		prev = EMPTY();;\
	};;\
	return STRUCT(){with}, prev, ok;;\
};;\
;;\
func (a STRUCT()) Delete(x K) (STRUCT(), T, bool) {;;\
	DO_SEARCH(a.data, ID(x), i, has);;\
	if !has {;;\
		return a, EMPTY(), false;\
	};;\
	without := make(SLICE(T), len(a.data)-1);;\
	copy(without[:i], a.data[:i]);;\
	copy(without[i:], a.data[i+1:]);;\
	return STRUCT(){without}, a.data[i], true;;\
};;\
;;\
func (a STRUCT()) Ascend(cb func(x T) bool) bool {;;\
	for _, x := range a.data {;;\
		if !cb(x) {;;\
			return false;;\
		};;\
	};;\
	return true;;\
};;\
;;\
func (a STRUCT()) AscendRange(x, y K, cb func(x T) bool) bool {;;\
	DO_SEARCH_RANGE(a.data, ID(x), 0, len(a.data), i, hasX);;\
	DO_SEARCH_RANGE(a.data, ID(y), i, len(a.data), j, hasY);;\
	for ; i < len(a.data) && i <= j; i++ {;;\
		if !cb(a.data[i]) {;;\
			return false;;\
		};;\
	};;\
	return true;;\
};;\
;;\
func (a STRUCT()) Reset() STRUCT() {;;\
	return STRUCT(){nil};;\
};;\
;;\
func (a STRUCT()) AppendTo(p SLICE(T)) SLICE(T) {;;\
	return append(p, a.data...);;\
};;\
;;\
func (a STRUCT()) Len() int {;;\
	return len(a.data);;\
};;\
;;\
func (a STRUCT()) Cap() int {;;\
	return cap(a.data);;\
};;\

#endif /* !_PPGO_STRUCT_ARRAY_ */
