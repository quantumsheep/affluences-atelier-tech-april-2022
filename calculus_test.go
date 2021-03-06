package main

import "testing"

func TestAddIntegers(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Add(1, 2) != 3")
	}
}

func TestAddStrings(t *testing.T) {
	if Add("a", "b") != "ab" {
		t.Error(`Add("a", "b") != "ab"`)
	}
}

func TestSubIntegers(t *testing.T) {
	if Sub(1, 2) != -1 {
		t.Error("Sub(1, 2) != -1")
	}
}

func TestMulIntegers(t *testing.T) {
	if Mul(1, 2) != 2 {
		t.Error("Mul(1, 2) != 2")
	}
}

func TestDivIntegers(t *testing.T) {
	if Div(6, 2) != 3 {
		t.Error("Div(6, 2) != 3")
	}
}

func TestModIntegers(t *testing.T) {
	if Mod(6, 2) != 0 {
		t.Error("Mod(6, 2) != 0")
	}
}

func TestPowIntegers(t *testing.T) {
	if Pow(2, 3) != 8 {
		t.Error("Pow(2, 3) != 8")
	}
}
