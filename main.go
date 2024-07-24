package main

import (
	"fmt"
	"tutorial/day01"
	"tutorial/lib"
)

func main() {
	investmentExample()

	for i := 0; i < 10; i++ {
		fmt.Println(lib.TestConditionals(i))
	}

	fmt.Println(`Testing switch...`)
	fmt.Println(lib.TestSwitch(3))
	fmt.Println()

	day01.ProfitCalculatorApplication()
}

func investmentExample() {
	var investmentAmount, expectedReturnRate float64
	const years int = 10

	fmt.Print(`Enter investment amount: `)
	fmt.Scan(&investmentAmount)

	expectedReturnRate = 5.5

	formattedFutureValue := fmt.Sprintf("Future Value: %.2f\n", lib.CalculateFutureValue(investmentAmount, years, expectedReturnRate))
	fmt.Print(formattedFutureValue)
}
