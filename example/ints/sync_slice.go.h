#include "ppgo/struct/sync_slice.h"

#define ID(a) a
#define LESS_OR_EQUAL(a, b) a <= b
#define GREATER(a, b) a > b
#define FUNC(a) a
#define STRUCT() SyncSlice
#define CTOR() NewSyncSlice
#define EMPTY() 0

package ints

MAKE_SYNC_SORTED_SLICE(int, int)
