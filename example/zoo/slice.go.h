#include "ppgo/struct/slice.h"

#define ID(a) a.Name
#define LESS_OR_EQUAL(a, b) a.Name <= b.Name
#define FUNC(a) a
#define STRUCT(a) Animals
#define CTOR() NewAnimals
#define EMPTY() Animal{}

package zoo

MAKE_SORTED_SLICE(Animal, string)
