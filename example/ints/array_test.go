package ints

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArrayUpsertFill(t *testing.T) {
	tuple := Array{}

	for i := 0; i < ArrayCapacity; i++ {
		tuple, _, _ = tuple.Upsert(i)
	}

	done := make(chan error)
	go func() {
		defer func() {
			if reason := recover(); reason != nil {
				done <- fmt.Errorf("panic: %s", reason)
			} else {
				close(done)
			}
		}()
		tuple, _, _ = tuple.Upsert(42)
	}()

	if err := <-done; err == nil {
		t.Errorf("Upsert() should panic after overfilling insertion attempt")
	}
}

func TestArrayUpsertDelete(t *testing.T) {
	tuple := Array{}

	tuple, _, _ = tuple.Upsert(1)
	tuple, _, _ = tuple.Upsert(3)
	tuple, _, _ = tuple.Upsert(2)

	if act, exp := tuple.AppendTo(nil), []int{1, 2, 3}; !reflect.DeepEqual(act, exp) {
		t.Errorf("after upserting 1,3,2 AppendTo() returns %v; want %v", act, exp)
	}

	tuple, _, _ = tuple.Delete(2)

	if act, exp := tuple.AppendTo(nil), []int{1, 3}; !reflect.DeepEqual(act, exp) {
		t.Errorf("after upserting 1,3,2 and deleting 2 AppendTo() returns %v; want %v", act, exp)
	}
}
