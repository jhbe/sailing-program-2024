package main

import "testing"

func Test_Next(t *testing.T) {
	pros := Pros{
		Names: []string{"A", "B", "C"},
	}

	if name := pros.Next(); name != "A" {
		t.Error("Expected A")
	}
	if name := pros.Next(); name != "B" {
		t.Error("Expected B")
	}
	if name := pros.Next(); name != "C" {
		t.Error("Expected C")
	}
	if name := pros.Next(); name != "A" {
		t.Error("Expected A")
	}
}
