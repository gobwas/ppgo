package ints

import (
	"fmt"
	"testing"
)

func TestSyncArrayGetAny(t *testing.T) {
	for _, test := range []struct {
		data   []int
		get    []int
		expect int
		ok     bool
	}{
		{
			data:   []int{0, 1, 2, 4},
			get:    []int{-2, -1, 0, 1},
			expect: 0,
			ok:     true,
		},
	} {
		t.Run("", func(t *testing.T) {
			s := NewSyncArray(0)
			for _, v := range test.data {
				s.Upsert(v)
			}
			v, ok := s.GetAny(iterator(test.get))
			if ok != test.ok || v != test.expect {
				t.Errorf("GetAny(%v) = %v, %v; want %v, %v", test.get, v, ok, test.expect, test.ok)
			}
		})
	}
}

func TestSyncArrayGetsertAny(t *testing.T) {
	for i, test := range []struct {
		data   []int
		get    []int
		insert int
		expect int
	}{
		{
			data:   []int{0, 1, 2, 4},
			get:    []int{-2, -1, 0, 1},
			insert: -42,
			expect: 0,
		},
		{
			data:   []int{0, 1, 2, 4},
			get:    []int{-4, -3, -2, -1},
			insert: 42,
			expect: 42,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			s := NewSyncArray(0)
			for _, v := range test.data {
				s.Upsert(v)
			}
			v := s.GetsertAnyFn(iterator(test.get), func() int {
				return test.insert
			})
			if v != test.expect {
				t.Errorf(
					"GetsertAnyFn(%v, => %v) = %v; want %v",
					test.get, test.insert, v, test.expect,
				)
			}
		})
	}
}

func iterator(data []int) func() (int, bool) {
	var i int
	return func() (int, bool) {
		if i >= len(data) {
			return 0, false
		}
		defer func() { i++ }()
		return data[i], true
	}
}
