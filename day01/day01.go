package day01

import "fmt"

func ProfitCalculatorApplication() {
	fmt.Println("Profit Calculator Application")

	revenue := getUserInput("Enter revenue: ")
	expenses := getUserInput("Enter expenses: ")
	taxRate := getUserInput("Enter tax rate: ")

	var ebt = earningsBeforeTax(revenue, expenses)

	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)

	var eat = earningsAfterTax(revenue, expenses, taxRate)

	fmt.Printf("Earnings After Tax (EAT): %.2f\n", eat)

	var npm = netProfitMargin(revenue, expenses, taxRate)

	fmt.Printf("Net Profit Margin (NPM): %.2f\n", npm)
}

func getUserInput(infoText string) float64 {
	var value float64
	fmt.Print(infoText)
	fmt.Scan(&value)
	return value
}

// Calculates the earnings before tax (EBT)
func earningsBeforeTax(revenue float64, expenses float64) float64 {
	return revenue - expenses
}

// Calculate the earnings after tax (Net Income)
func earningsAfterTax(revenue float64, expenses float64, taxRate float64) float64 {
	return earningsBeforeTax(revenue, expenses) * (1.0 - taxRate/100)
}

// Calculate the ration between EBT and profit (Net Profit Margin)
func netProfitMargin(revenue float64, expenses float64, taxRate float64) float64 {
	return earningsBeforeTax(revenue, expenses) / earningsAfterTax(revenue, expenses, taxRate)
}
