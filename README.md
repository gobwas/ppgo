(P)re(P)rocessing (GO)
======================

## Status: ßeta

### Benchmarks

```
BenchmarkSort/ppgo_10-4                  5000000               259 ns/op               0 B/op          0 allocs/op
BenchmarkSort/ppgo_100-4                  300000              4758 ns/op               0 B/op          0 allocs/op
BenchmarkSort/ppgo_1000-4                  30000             51842 ns/op               0 B/op          0 allocs/op
BenchmarkSort/golang_10-4                1000000              1803 ns/op              32 B/op          1 allocs/op
BenchmarkSort/golang_100-4                200000             10110 ns/op              32 B/op          1 allocs/op
BenchmarkSort/golang_1000-4                10000            122353 ns/op              32 B/op          1 allocs/op
```

### Example

Somewhere inside your project `strings/strings.go`:
```go
package strings

//go:generate ppgo
```

Put the header file `strings/strings.go.h`:
```cpp
#include "ppgo_sort.h"

#define ID(a) a
#define LESS_OR_EQUAL(a, b) a <= b
#define GREATER(a, b) a > b
#define FUNC(a) MyPrefix##a

package strings

MAKE_SORT(string, string)
```

Run generation: 
```shell
go generate ./strings
```

Enjoy! :)

### More Examples

Please see `./example` folder to checkout current possibilities of `ppgo`.
