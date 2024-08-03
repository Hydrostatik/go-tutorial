package day03

import (
	"testing"
)

func TestDay03(t *testing.T) {
	t.Run("Correctly calculates the Area", func(t *testing.T) {
		result := Rectangle{
			Width:  4,
			Length: 5,
		}.Area()

		if result != 20 {
			t.Errorf("Result was incorrect, got %f, want: %f.", result, 20.0)
		}
	})
	t.Run("Correctly calculates the Perimeter", func(t *testing.T) {
		result := Rectangle{
			Width:  4,
			Length: 5,
		}.Perimeter()

		if result != 18 {
			t.Errorf("Result was incorrect, got %f, want: %f.", result, 18.0)
		}
	})
}
