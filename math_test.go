package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(15, 1)

	if total != 30 {
		t.Errorf("Invalid sum result: Result: %d. Expected: %d", total, 30)
	}
}
