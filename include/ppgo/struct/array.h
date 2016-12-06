#include "ppgo/algorithm/sort.h"

/**
 * This file contains an implementation of immutable sorted array.
 */

#ifndef _PPGO_STRUCT_ARRAY_
#define	_PPGO_STRUCT_ARRAY_

#define MAKE_ARRAY(T, K);;\
type STRUCT() struct {;;\
	data SLICE(T);;\
};;\
;;\
func (a STRUCT()) Has(x K) bool {;;\
	DO_SEARCH(a.data, x, i, ok);;\
	return ok;;\
};;\
;;\
func (a STRUCT()) Get(x K) (T, bool) {;;\
	DO_SEARCH(a.data, x, i, ok);;\
	if !ok {;;\
		return EMPTY(), false;;\
	};;\
	return a.data[i], true;;\
};;\
;;\
func (a STRUCT()) Upsert(x T) (cp STRUCT(), prev T) {;;\
	var with SLICE(T);;\
	DO_SEARCH(a.data, ID(x), i, has);;\
	if has {;;\
		with = make(SLICE(T), len(a.data));;\
		copy(with, a.data);;\
		with[i], prev = x, a.data[i];;\
	} else {;;\
		with = make(SLICE(T), len(a.data)+1);;\
		copy(with[:i], a.data[:i]);;\
		copy(with[i+1:], a.data[i:]);;\
		with[i] = x;;\
	};;\
	return STRUCT(){with}, prev;;\
};;\
;;\
func (a STRUCT()) Delete(x K) (STRUCT(), T, bool) {;;\
	DO_SEARCH(a.data, x, i, has);;\
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
	DO_SEARCH_RANGE(a.data, x, 0, len(a.data), i, hasX);;\
	DO_SEARCH_RANGE(a.data, y, i, len(a.data), j, hasY);;\
	for ; i < len(a.data) && i <= j; i++ {;;\
		if !cb(a.data[i]) {;;\
			return false;;\
		};;\
	};;\
	return true;;\
};;\
;;\
func (a STRUCT()) Reset() (STRUCT()) {;;\
	return STRUCT(){nil};;\
};;\
;;\
func (a STRUCT()) Len() int {;;\
	return len(a.data);;\
};;\

#endif /* !_PPGO_STRUCT_ARRAY_ */
