#include "ppgo/util.h"

/**
 * This file contains an implementation of d-ary heap.
 */

#ifndef _PPGO_STRUCT_HEAP_
#define _PPGO_STRUCT_HEAP_

#define MAKE_HEAP(T);;\
type STRUCT() struct {;;\
	d     int;;\
	data  []T;;\
};;\
;;\
func CTOR()(d int) *STRUCT() {;;\
	return &STRUCT() {;;\
		d: d,;;\
	};;\
};;\
;;\
func CONCAT(CTOR(), FromSlice)(data []T, d int) *STRUCT() {;;\
	h := &STRUCT(){;;\
		d:    d,;;\
		data: data,;;\
	};;\
	_HEAPIFY(_HEAP_SWAP);;\
	return h;;\
};;\
;;\
func (h *STRUCT()) Top() T {;;\
	if len(h.data) == 0 {;;\
		return EMPTY();;\
	};;\
	return h.data[0];;\
};;\
;;\
func (h *STRUCT()) Pop() T {;;\
	n := len(h.data);;\
	ret := h.data[0];;\
	SWAP(h.data, 0, n-1);;\
	h.data[n-1] = EMPTY();;\
	h.data = h.data[:n-1];;\
	h.siftDown(0);;\
	return ret;;\
};;\
;;\
func (h *STRUCT()) Push(x T) {;;\
	_PUSH_BACK(x, i);;\
	h.siftUp(i);;\
};;\
;;\
func (h *STRUCT()) Heapify() {;;\
	_HEAPIFY(_HEAP_SWAP);;\
};;\
;;\
func (h *STRUCT()) Slice() []T {;;\
	cp := *h;;\
	cp.data = make([]T, len(h.data));;\
	copy(cp.data, h.data);;\
	ret := cp.data;;\
	n := len(cp.data) - 1;;\
	for i := 0; i < n; i++ {;;\
		cp.data = cp.data[1:];;\
		cp.siftDown(0);;\
	};;\
	return ret;;\
};;\
;;\
func (h *STRUCT()) Data() []T { return h.data };;\
func (h *STRUCT()) Len() int { return len(h.data) } ;;\
func (h *STRUCT()) Empty() bool { return len(h.data) == 0 };;\
;;\
func (h STRUCT()) siftDown(root int) {;;\
	_SIFT_DOWN(root, _HEAP_SWAP, LESS_OR_EQUAL);;\
};;\
;;\
func (h STRUCT()) siftUp(root int) {;;\
	_SIFT_UP(root, _HEAP_SWAP, LESS_OR_EQUAL);;\
};;\
;;\
func (h STRUCT()) siftTop(root int) {;;\
	_SIFT_TOP(root, _HEAP_SWAP);;\
};;\

#define _HEAP_SWAP(i, j)\
	SWAP(h.data, i, j)\

#define _HEAPIFY(SWAP_FN)\
	for i := len(h.data)/h.d - 1; i >= 0; i-- {;;\
		h.siftDown(i);;\
	}\

#define _PUSH_BACK(x, n)\
	n := len(h.data);;\
	if cap(h.data) == len(h.data) {;;\
		h.data = append(h.data, x);;\
	} else {;;\
		h.data = h.data[:n+1];;\
		h.data[n] = x;;\
	}\

#define _SIFT_DOWN(root, SWAP_FN, LESS_OR_EQUAL_FN)\
	for {;;\
		min := root;;\
		for i := 1; i <= h.d; i++ {;;\
			child := h.d*root + i;;\
			if child >= len(h.data) { >>> out of bounds;;\
				break;;\
			};;\
			if !(LESS_OR_EQUAL_FN(h.data[min], h.data[child])) {;;\
				min = child;;\
			};;\
		};;\
		if min == root {;;\
			return;;\
		};;\
		SWAP_FN(root, min);;\
		root = min;;\
	}\


#define _SIFT_UP(root, SWAP_FN, LESS_OR_EQUAL_FN)\
	for root > 0 {;;\
		parent := (root - 1) / h.d;;\
		if !(LESS_OR_EQUAL_FN(h.data[root], h.data[parent])) {;;\
			return;;\
		};;\
		SWAP_FN(parent, root);;\
		root = parent;;\
	}\


#define _SIFT_TOP(root, SWAP_FN)\
	for root > 0 {;;\
		parent := (root - 1) / h.d;;\
		SWAP_FN(parent, root);;\
		root = parent;;\
	}\

#endif /* !_PPGO_STRUCT_HEAP_ */
