#include "ppgo/struct/heap.h"

#define LESS_OR_EQUAL(a, b) a <= b
#define STRUCT(a) Heap
#define CTOR() NewHeap
#define EMPTY() 0

package ints

MAKE_HEAP(int)
