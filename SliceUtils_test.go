package main

import (
	"testing"
)

func TestFilter(t *testing.T) {
	var slice = []string{"1", "2", "22"}
	var filtered = Filter(slice, func(elem string) bool {
		return len(elem)%2 == 0
	})
	AssertEqualsInt(t, 1, len(filtered))
	AssertEqualsString(t, "22", filtered[0])
}

func TestMapString(t *testing.T) {
	var slice = []string{"1", "2"}
	var filtered = MapString(slice, func(elem string) string {
		return elem + elem
	})
	AssertEqualsInt(t, 2, len(filtered))
	AssertEqualsString(t, "11", filtered[0])
	AssertEqualsString(t, "22", filtered[1])
}