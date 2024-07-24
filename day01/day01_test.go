package day01

import (
	"testing"
)

func TestDay01(t *testing.T) {
	t.Run("EBT calculated correctly", func(t *testing.T) {
		result := earningsBeforeTax(4, 3)
		if result != 1 {
			t.Errorf("Result was incorrect, got: %f, want: %f.", result, 1.0)
		}
	})
	t.Run("EAT calculated correctly", func(t *testing.T) {
		result := earningsAfterTax(4, 3, 50)
		if result != 0.5 {
			t.Errorf("Result was incorrect, got: %f, want: %f.", result, 0.5)
		}
	})
	t.Run("NPM calculated correctly", func(t *testing.T) {
		result := netProfitMargin(4, 3, 50)
		if result != 2 {
			t.Errorf("Result was incorrect, got: %f, want: %f.", result, 2.0)
		}
	})
}
