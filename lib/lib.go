package lib

import "math"

// Takes in an investment amount, years invested and expected return (percent)
// and outputs the future value
func CalculateFutureValue(investmentAmount float64, years int, expectedReturn float64) float64 {
	return investmentAmount * math.Pow(float64(1+expectedReturn/100), float64(years))
}

// Test out how conditionals work in go
func TestConditionals(x int) string {
	if x == 1 {
		return "Your number is one"
	} else if x == 0 {
		return "Story of your life"
	}

	return "That's a weird number..."
}

// Test out how switch statements work in go
func TestSwitch(x int) string {
	switch x {
	case 1:
		return "Your number is one"
	case 0:
		return "Story of your life"
	default:
		return "That's a weird number..."
	}
}
