#include "sync_array.h"

#define ID(a) a
#define LESS_OR_EQUAL(a, b) a <= b
#define GREATER(a, b) a > b
#define FUNC(a) a
#define STRUCT() SyncArray
#define CTOR() NewSyncArray
#define EMPTY() 0

package ints

MAKE_ARRAY(int, int)
