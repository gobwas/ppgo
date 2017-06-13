#include "ppgo/struct/tuple.h"

#define ID(a) a
#define LESS_OR_EQUAL(a, b) a <= b
#define STRUCT(a) Tuple
#define VAR(a) CONCAT(Tuple, a)
#define EMPTY() 0

package ints

MAKE_TUPLE(8, int, int)
