package main

import "fmt"

func main() {

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax rate: ")

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	printResult(ebt, profit, ratio)

}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func printResult(ebt, profit, ratio float64) {
	fmt.Printf("%.1f", ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}