package main

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(5, 5); got != 10 {
		t.Errorf("Add(5,5) = %d; want 10", got)
	}
}
