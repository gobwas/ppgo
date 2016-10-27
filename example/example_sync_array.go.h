#include "sync_array.h"

#define ID(a) a
#define LESS_OR_EQUAL(a, b) a <= b
#define GREATER(a, b) a > b
#define FUNC(a) uint##a
#define EMPTY() 0
#define STRUCT(a) uint##a
#define CTOR(a) newUint##a

package ppgo_test

MAKE_ARRAY(uint, uint)

