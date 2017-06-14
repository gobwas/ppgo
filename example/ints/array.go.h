#include "ppgo/struct/array.h"

#define ID(a) a
#define LESS_OR_EQUAL(a, b) a <= b
#define STRUCT(a) Array
#define VAR(a) CONCAT(Array, a)
#define EMPTY() 0

package ints

MAKE_SORTED_ARRAY(8, int, int)
