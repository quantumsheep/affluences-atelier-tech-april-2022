package main

import "testing"

func TestAddIntegers(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Add(1, 2) != 3")
	}
}
