#include "ppgo/util.h"
#include "ppgo/struct/heap.h"

/**
 * This file contains an implementation of d-ary heap.
 */

#ifndef _PPGO_STRUCT_INDEXED_HEAP_
#define _PPGO_STRUCT_INDEXED_HEAP_

#define MAKE_INDEXED_HEAP(T, W);;\
	_MAKE_INDEXED_HEAP(T, W, CONCAT(record, STRUCT()));;\

#define _MAKE_INDEXED_HEAP(T, W, R);;\
type R struct {;;\
	x T;;\
	w W;;\
};;\
;;\
;;\
type STRUCT() struct {;;\
	d     int;;\
	data  []R;;\
	index map[T]int;;\
};;\
;;\
func CTOR()(d int) *STRUCT() {;;\
	return &STRUCT() {;;\
		d: d,;;\
		index: make(map[T]int),;;\
	};;\
};;\
;;\
func CONCAT(CTOR(), FromSlice)(data []T, d int) *STRUCT() {;;\
	records := make([]R, len(data));;\
	for i, x := range data {;;\
		records[i] = R{x:x};;\
	};;\
	h := &STRUCT(){;;\
		d:     d,;;\
		data:  records,;;\
		index: make(map[T]int),;;\
	};;\
	_HEAPIFY(_INDEXED_HEAP_SWAP);;\
	return h;;\
};;\
;;\
func (h *STRUCT()) Top() T {;;\
	if len(h.data) == 0 {;;\
		return EMPTY();;\
	};;\
	return h.data[0].x;;\
};;\
;;\
func (h *STRUCT()) Pop() T {;;\
	n := len(h.data);;\
	ret := h.data[0].x;;\
\
	_INDEXED_HEAP_SWAP(0, n-1);;\
\
	h.data[n-1] = R{};;\
	h.data = h.data[:n-1];;\
	delete(h.index, ret);;\
\
	h.siftDown(0);;\
\
	return ret;;\
};;\
;;\
func (h *STRUCT()) Slice() []T {;;\
	cp := *h;;\
	cp.data = make([]R, len(h.data));;\
	copy(cp.data, h.data);;\
	cp.index = make(map[T]int) >>> prevent reordering original index;;;\
	ret := make([]T, len(cp.data));;\
	n := len(cp.data);;\
	for i := 0; i < n; i++ {;;\
		ret[i] = cp.data[0].x;;\
		cp.data = cp.data[1:];;\
		cp.siftDown(0);;\
	};;\
	return ret;;\
};;\
;;\
func (h *STRUCT()) Len() int { return len(h.data) };;\
func (h *STRUCT()) Empty() bool { return len(h.data) == 0 };;\
;;\
func (h *STRUCT()) Push(x T, w W) {;;\
	_, ok := h.index[x];;\
	if ok {;;\
		panic("could not push value that is already present in heap");;\
	};;\
	r := R{x, w};;\
	_PUSH_BACK(r, i);;\
	h.index[x] = i;;\
	h.siftUp(i);;\
};;\
;;\
func (h *STRUCT()) Heapify() {;;\
	_HEAPIFY(_INDEXED_HEAP_SWAP);;\
};;\
;;\
func (h *STRUCT()) WithPriority(x T, fn func(W) W) {;;\
	i, ok := h.index[x];;\
	if !ok {;;\
		panic("could not update value that is not present in heap");;\
	};;\
	h.update(i, R{x, fn(h.data[i].w)});;\
};;\
;;\
func (h *STRUCT()) ChangePriority(x T, w W) {;;\
	i, ok := h.index[x];;\
	if !ok {;;\
		panic("could not update value that is not present in heap");;\
	};;\
	h.update(i, R{x, w});;\
};;\
;;\
func (h *STRUCT()) Compare(a, b T) int {;;\
	var i, j int;;\
	i, ok := h.index[a];;\
	if ok {;;\
		j, ok = h.index[b];;\
	};;\
	if !ok {;;\
		panic("comparing record that not in heap");;\
	};;\
	return COMPARE(h.data[i].w, h.data[j].w);;\
};;\
;;\
func (h *STRUCT()) Remove(x T) {;;\
	i, ok := h.index[x];;\
	if !ok {;;\
		return;;\
	};;\
	h.siftTop(i);;\
	h.Pop();;\
};;\
;;\
func (h *STRUCT()) update(i int, r R) {;;\
	prev := h.data[i];;\
	h.data[i] = r;;\
	if !(LESS_OR_EQUAL(r.w, prev.w)) {;;\
		h.siftDown(i);;\
	} else {;;\
		h.siftUp(i);;\
	};;\
};;\
;;\
func (h STRUCT()) siftDown(root int) {;;\
	_SIFT_DOWN(root, _INDEXED_HEAP_SWAP, _INDEXED_HEAP_LESS_OR_EQUAL);;\
};;\
;;\
func (h STRUCT()) siftUp(root int) {;;\
	_SIFT_UP(root, _INDEXED_HEAP_SWAP, _INDEXED_HEAP_LESS_OR_EQUAL);;\
};;\
;;\
func (h STRUCT()) siftTop(root int) {;;\
	_SIFT_TOP(root, _INDEXED_HEAP_SWAP);;\
};;\

#define _INDEXED_HEAP_LESS_OR_EQUAL(a, b)\
	LESS_OR_EQUAL(a.w, b.w)\

#define _INDEXED_HEAP_SWAP(i, j)\
	a, b := h.data[i], h.data[j];;\
	h.index[a.x], h.index[b.x] = j, i;;\
	SWAP(h.data, i, j)\

#endif /* !_PPGO_STRUCT_INDEXED_HEAP_ */
