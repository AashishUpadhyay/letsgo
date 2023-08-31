package main

import (
	"strconv"
	"testing"
)

func TestMenuOptions(t *testing.T) {
	// arrange
	want := true
	choice := "2"
	itemname := "Coffee"

	// act
	got := menu_options(choice, itemname)

	// assert
	if got != want {
		t.Errorf("got %q, wanted %q", strconv.FormatBool(got), strconv.FormatBool(want))
	}
}
