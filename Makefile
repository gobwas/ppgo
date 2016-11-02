TEMPLATES = $(wildcard $(PWD)/example/ints/*.go.h)
GENERATED = $(wildcard $(PWD)/example/ints/*_gen.go $(PWD)/*_gen_test.go)

BENCH ?= .

.PHONY: generate

all:

clean:
	for file in $(GENERATED); do [ -f $$file ] && rm $$file; done

bin/ppgo:
	go build -o bin/ppgo

generate: clean bin/ppgo
	PATH=$$PWD/bin:$$PATH go generate ./example/ints/...
	
test: 
	go test -v -cover ./example/ints/...

bench:
	go test -run=none -bench=$(BENCH) -benchmem ./example/ints/...
