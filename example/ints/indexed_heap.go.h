#include "ppgo/struct/indexed_heap.h"

#define LESS_OR_EQUAL(a, b) a <= b
#define STRUCT(a) IndexedHeap
#define CTOR() NewIndexedHeap
#define EMPTY() 0
#define COMPARE(a, b) a - b

package ints

MAKE_INDEXED_HEAP(int, int)
