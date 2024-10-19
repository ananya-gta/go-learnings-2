package main

import "fmt"
import "os"
import "errors"

func writeProfit(ebt, profit, ratio float64) {
	profits := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("profit.txt", []byte(profits), 0644)
}

func main() {

	revenue, err1 := getUserInput("Revenue: ")
	expenses, err2 := getUserInput("Expenses: ")
	taxRate, err3 := getUserInput("Tax rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("ERRRRRRROOOORRRR")
		fmt.Print(err1)
	}

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	printResult(ebt, profit, ratio)
	writeProfit(ebt, profit, ratio)

}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("invalid user input")
	}

	return userInput, nil
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func printResult(ebt, profit, ratio float64) {
	fmt.Println("Below are the following results:")
	fmt.Printf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
}
