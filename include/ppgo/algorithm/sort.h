#include "ppgo/util.h"

#ifndef _PPGO_ALGORITHM_SORT_
#define	_PPGO_ALGORITHM_SORT_

#define MAKE_SORT(T, K);;\
	func FUNC(QuickSort)(data SLICE(T), lo, hi int) {;;\
		if lo >= hi {;;\
			return;;\
		};;\
		DO_PARTITION(data, lo, hi, p);;\
		FUNC(QuickSort)(data, lo, p);;\
		FUNC(QuickSort)(data, p+1, hi);;\
	};;;;\
	func FUNC(InsertionSort)(data SLICE(T), l, r int) {;;\
		DO_INSERTION_SORT(data, l, r)\
	};;;;\
	func FUNC(Sort)(data SLICE(T), l, r int) {;;\
		if r-l > 12 {;;\
			FUNC(QuickSort)(data, l, r);;\
			return;;\
		};;\
		DO_INSERTION_SORT(data, l, r);;\
	};;;;\
	func FUNC(Search)(data SLICE(T), key K) (int, bool) {;;\
		DO_SEARCH(data, key, i, ok);;\
		return i, ok;;\
	};;;;\

#define DO_SWAP(DATA, A, B)\
	DATA[A], DATA[B] = DATA[B], DATA[A]\

#define DO_PARTITION(DATA, L, R, PIVOT)\
	>>> Quick sort partition algorithm.;;\
	var PIVOT int;;\
	{;;\
		>>> Let x be a pivot;;\
		x := DATA[L];;\
		PIVOT = L;;\
		for i := L + 1; i < R; i++ {;;\
			if LESS_OR_EQUAL(DATA[i], x) {;;\
				PIVOT++;;\
				DO_SWAP(DATA, PIVOT, i);;\
			};;\
		};;\
		DO_SWAP(DATA, PIVOT, L);;\
	}\

#define DO_INSERTION_SORT(DATA, L, R)\
	>>> Insertion sort algorithm.;;\
	for i := L + 1;; i < R;; i++ {;;\
		for j := i;; j > L && !(LESS_OR_EQUAL(DATA[j-1], DATA[j]));; j-- {;;\
			DO_SWAP(DATA, j, j-1);;\
		};;\
	}\

#define DO_SEARCH_RANGE(DATA, KEY, LEFT, RIGHT, RESULT, OK)\
	>>> Binary search algorithm.;;\
	var OK bool;;\
	var RESULT int;;\
	{;;\
		l := LEFT;;\
		r := RIGHT;;\
		for !OK && l < r {;;\
			m := l + (r-l)/2;;\
			switch {;;\
			case ID(DATA[m]) == KEY:;;\
				OK = true;;\
				r = m;;\
			case ID(DATA[m]) < KEY:;;\
				l = m + 1;;\
			case ID(DATA[m]) > KEY:;;\
				r = m;;\
			};;\
		};;\
		RESULT = r;;\
		_ = RESULT >>> in case when RESULT not being used;;\
	}\
	

#define DO_SEARCH(DATA, KEY, RESULT, OK)\
	DO_SEARCH_RANGE(DATA, KEY, 0, len(DATA), RESULT, OK)\

#define DO_SEARCH_SHORT(DATA, KEY, RIGHT)\
	DO_SEARCH(DATA, KEY, RIGHT, CONCAT(ok, __COUNTER__))\

#endif /* !_PPGO_ALGORITHM_SORT_ */
