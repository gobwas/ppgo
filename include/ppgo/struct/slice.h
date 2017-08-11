#include "ppgo/algorithm/sort.h"

/**
 * This file contains an implementation of immutable sorted slice.
 */

#ifndef _PPGO_STRUCT_SORTED_SLICE_
#define	_PPGO_STRUCT_SORTED_SLICE_

#define MAKE_SORTED_SLICE(T, K);;\
type STRUCT() struct {;;\
	Data SLICE(T);;\
};;\
;;\
>>> CTOR() creates STRUCT() with underlying data.;;\
>>> Note that data is not copied and used by reference.;;\
func CTOR()(data SLICE(T)) STRUCT() {;;\
	PRIVATE_FUNC(STRUCT(), SortSource)(data, 0, len(data));;\
	return STRUCT(){Data: data};;\
};;\
;;\
>>> PRIVATE_FUNC(STRUCT(), SortSource) sorts data for further use inside STRUCT().;;\
func PRIVATE_FUNC(STRUCT(), SortSource)(data SLICE(T), lo, hi int) {\
	MK_SORT(PRIVATE_FUNC(STRUCT(), SortSource), data, lo, hi);;\
};;\
;;\
func (a STRUCT()) Has(x K) bool {;;\
	DO_SEARCH(a.Data, x, i, ok);;\
	return ok;;\
};;\
;;\
func (a STRUCT()) Get(x K) (T, bool) {;;\
	DO_SEARCH(a.Data, x, i, ok);;\
	if !ok {;;\
		return EMPTY(), false;;\
	};;\
	return a.Data[i], true;;\
};;\
;;\
>>> Upsert inserts item x into array or updates existing one.;;\
>>> It returns copy of STRUCT(), previous item (if were present) and a boolean;;\
>>> flag that reports about previous item replacement. This flag is useful for;;\
>>> non-pointer item types such as numbers or struct values.;;\
func (a STRUCT()) Upsert(x T) (cp STRUCT(), prev T, swapped bool) {;;\
	var with SLICE(T);;\
	DO_SEARCH(a.Data, ID(x), i, has);;\
	if has {;;\
		with = make(SLICE(T), len(a.Data));;\
		copy(with, a.Data);;\
		with[i], prev = x, a.Data[i];;\
		swapped = true;;\
	} else {;;\
		with = make(SLICE(T), len(a.Data)+1);;\
		copy(with[:i], a.Data[:i]);;\
		copy(with[i+1:], a.Data[i:]);;\
		with[i] = x;;\
		prev = EMPTY();;\
	};;\
	return STRUCT(){with}, prev, swapped;;\
};;\
;;\
func (a STRUCT()) Delete(x K) (STRUCT(), T, bool) {;;\
	DO_SEARCH(a.Data, x, i, has);;\
	if !has {;;\
		return a, EMPTY(), false;\
	};;\
	without := make(SLICE(T), len(a.Data)-1);;\
	copy(without[:i], a.Data[:i]);;\
	copy(without[i:], a.Data[i+1:]);;\
	return STRUCT(){without}, a.Data[i], true;;\
};;\
;;\
func (a STRUCT()) Ascend(cb func(x T) bool) bool {;;\
	for _, x := range a.Data {;;\
		if !cb(x) {;;\
			return false;;\
		};;\
	};;\
	return true;;\
};;\
;;\
func (a STRUCT()) AscendRange(x, y K, cb func(x T) bool) bool {;;\
	DO_SEARCH_RANGE(a.Data, x, 0, len(a.Data), i, hasX);;\
	DO_SEARCH_RANGE(a.Data, y, i, len(a.Data), j, hasY);;\
	for ; i < len(a.Data) && i <= j; i++ {;;\
		if !cb(a.Data[i]) {;;\
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
	return append(p, a.Data...);;\
};;\
;;\

#endif /* !_PPGO_STRUCT_SORTED_SLICE_ */
