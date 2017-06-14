#include "ppgo/struct/slice.h"

#define ID(a) a
#define LESS_OR_EQUAL(a, b) a <= b
#define FUNC(a) a
#define STRUCT(a) SortedSlice
#define CTOR() NewSortedSlice
#define EMPTY() 0

package ints

MAKE_SORTED_SLICE(int, int)
