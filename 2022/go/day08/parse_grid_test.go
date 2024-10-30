package main

import (
	"reflect"
	"testing"
)

func TestGridParse(t *testing.T) {
	got := ParseGrid("test.txt")
	want := [][]int{{3, 0, 3, 7, 3}, {2, 5, 5, 1, 2}, {6, 5, 3, 3, 2}, {3, 3, 5, 4, 9}, {3, 5, 3, 9, 0}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
