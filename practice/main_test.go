package main

import (
	"testing"
)

func TestCalculatePostive(t *testing.T) {
	got := calculate("10.0", "5.5")
	var want float64 = 15.5

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestCalcalateNegivite(t *testing.T) {
	got := calculate("100.0", "-110.5")
	var want float64 = -10.5

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
