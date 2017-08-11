package zoo

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortedSliceNew(t *testing.T) {
	for _, test := range []struct {
		in  []Animal
		exp []Animal
	}{
		{
			in:  animals("crocodile", "bee", "tiger"),
			exp: animals("bee", "crocodile", "tiger"),
		},
	} {
		t.Run("", func(t *testing.T) {
			s := NewSortedSlice(test.in)
			if !reflect.DeepEqual(s.data, test.exp) {
				t.Errorf("NewSortedSlice = %+v; want %v", s.data, test.exp)
			}
		})
	}
}

func TestSortedSliceGet(t *testing.T) {
	for _, test := range []struct {
		in  []Animal
		get string
		exp bool
	}{
		{
			in:  animals("crocodile", "bee", "tiger"),
			get: "wolf",
			exp: false,
		},
		{
			in:  animals("crocodile", "bee", "tiger"),
			get: "crocodile",
			exp: true,
		},
	} {
		t.Run("", func(t *testing.T) {
			s := NewSortedSlice(test.in)
			if _, ok := s.Get(test.get); ok != test.exp {
				t.Errorf("Get(%s) from %s: %t; want %t", test.get, s.data, ok, test.exp)
			}
		})
	}
}

func TestSortedSliceUpsert(t *testing.T) {
	for _, test := range []struct {
		in      []Animal
		upsert  string
		expSwap bool
		expData []Animal
	}{
		{
			in:      animals("crocodile", "bee", "tiger"),
			upsert:  "wolf",
			expSwap: false,
			expData: sortedAnimals("crocodile", "bee", "tiger", "wolf"),
		},
		{
			in:      animals("crocodile", "bee", "tiger"),
			upsert:  "tiger",
			expSwap: true,
			expData: sortedAnimals("crocodile", "bee", "tiger"),
		},
	} {
		t.Run("", func(t *testing.T) {
			s := NewSortedSlice(test.in)
			upsert := Animal{Name: test.upsert}
			cp, _, swapped := s.Upsert(upsert)
			if swapped != test.expSwap {
				t.Errorf(
					"Upsert(%s) to %s: swapped %t; want %t",
					test.upsert, s.data, swapped, test.expSwap,
				)
			}
			if !reflect.DeepEqual(cp.data, test.expData) {
				t.Errorf("Upsert(%s) = %+v; want %v", test.upsert, cp.data, test.expData)
			}
		})
	}
}

func TestSortedSliceDelete(t *testing.T) {
	for _, test := range []struct {
		in        []Animal
		del       string
		expRemove bool
		expData   []Animal
	}{
		{
			in:        animals("crocodile", "bee", "tiger"),
			del:       "wolf",
			expRemove: false,
			expData:   sortedAnimals("crocodile", "bee", "tiger"),
		},
		{
			in:        animals("crocodile", "bee", "tiger"),
			del:       "tiger",
			expRemove: true,
			expData:   sortedAnimals("crocodile", "bee"),
		},
	} {
		t.Run("", func(t *testing.T) {
			s := NewSortedSlice(test.in)
			cp, _, removed := s.Delete(test.del)
			if removed != test.expRemove {
				t.Errorf(
					"Delete(%s) to %s: removed %t; want %t",
					test.del, s.data, removed, test.expRemove,
				)
			}
			if !reflect.DeepEqual(cp.data, test.expData) {
				t.Errorf("Delete(%s) = %+v; want %v", test.del, cp.data, test.expData)
			}
		})
	}
}

func sortedAnimals(s ...string) []Animal {
	sort.Strings(s)
	return animals(s...)
}

func animals(s ...string) []Animal {
	as := make([]Animal, len(s))
	for i, name := range s {
		as[i] = Animal{Name: name}
	}
	return as
}
