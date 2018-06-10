package main

import "testing"

func TestAdd(t *testing.T) {
	v := Add(3, 4)
	if v != 7 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", v, 7)
	}
}
