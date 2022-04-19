package main

import "testing"

func TestAddIntegers(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Add(1, 2) != 3")
	}
}

func TestSubIntegers(t *testing.T) {
	if Sub(1, 2) != -1 {
		t.Error("Sub(1, 2) != -1")
	}
}
