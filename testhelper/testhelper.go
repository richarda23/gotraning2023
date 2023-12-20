package testhelper

import (
	"math"
	"slices"
	"testing"
)

func AssertInteger(t testing.TB, wanted, got int) {
	t.Helper()
	if wanted != got {
		t.Errorf("Wanted %d but got %d", wanted, got)
	}
}

func AssertFloat(t testing.TB, wanted, got float32) {
	t.Helper()
	if math.Abs(float64(wanted)-float64(got)) > 0.1 {
		t.Errorf("Wanted %.2f but got %.2f", wanted, got)
	}
}

func AssertCorrectMessage(t testing.TB, got, wanted string) {
	t.Helper()
	if got != wanted {
		t.Errorf("wanted %q but got %q", wanted, got)
	}
}

func AssertArrayIntEqual(t testing.TB, got, wanted []int) {
	if !slices.Equal(got, wanted) {
		t.Errorf("expected %v but got %v", wanted, got)
	}
}

func AssertArrayStringEqual(t testing.TB, got, wanted []string) {
	if !slices.Equal(got, wanted) {
		t.Errorf("expected %v but got %v", wanted, got)
	}
}