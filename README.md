(P)re(P)rocessing (GO)
======================

## Status: ßeta

Example:

```
#include "sort.h"

#define ID(a) a
#define LESS_OR_EQUAL(a, b) a <= b
#define GREATER(a, b) a > b
#define FUNC(a) uint##a

package radix_test

MAKE_SORT(uint, uint)
```

Makefile:

```
TEMPLATES = $(wildcard $(PWD)/*.go.h)
GENERATED = $(wildcard $(PWD)/*_gen.go $(PWD)/*_gen_test.go)

.PHONY: generate

clean:
	for file in $(GENERATED); do [ -f $$file ] && rm $$file; done

generate:
	for tmpl in $(TEMPLATES); do \
		name=`basename $$tmpl .h`; \
		base=`basename $$name .go` \
		output="$${base}_gen.go"; \
		if [[ "$$base" =~ "_test" ]]; then \
			base=`basename $$base _test`; \
			output="$${base}_gen_test.go"; \
		fi; \
		tmp="$${output}.tmp"; \
	   	cc -Iinclude -E -P $$tmpl \
			| sed -E -e 's/>>>/\/\//g' \
			| sed -e $$'s/;;/\\\n/g' \
		   	> $$tmp; \
		gofmt $$tmp > $$output; \
		rm -f $$tmp; \
	done;
```
