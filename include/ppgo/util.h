#ifndef _PPGO_UTIL_
#define	_PPGO_UTIL_

#define SLICE(a) []a
#define ARRAY(n, a) [n]a 

#define _CONCAT(a, b) a ## b
#define CONCAT(a, b) _CONCAT(a, b)

#define SWAP(data, a, b)\
	data[a], data[b] = data[b], data[a]\

#endif /* !_PPGO_UTIL_ */
