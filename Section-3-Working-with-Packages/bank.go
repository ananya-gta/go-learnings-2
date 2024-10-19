package main

import (
	"example.com/splitting_code_across_different_packages/fileOps"
	"fmt"
	
)

func main() {
	var accountBalance, err = fileOps.ReadFloatValueFromFile("balance.txt", 1000)
	if err != nil {
		fmt.Println("ERRRRRRRORR")
		fmt.Println(err)
		panic("Application exited.")
	}
	fmt.Println("Welcome to Go Bank!")
	startBankApplication(accountBalance)
}

func startBankApplication(accountBalance float64) {
	for {
		printChoices()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your current balance is: ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Minimum deposit amount is Rs. 1")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your current balance is: ", accountBalance)
			fileOps.WriteIntoFile(accountBalance, "balance.txt")
		case 3:
			fmt.Print("Enter withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Print("Minimum withdar amount should be Rs. 100")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Print("Withdrawal amount should not be more than current balance. Your current balance is: ", accountBalance)
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Your current balance is: ", accountBalance)
			fileOps.WriteIntoFile(accountBalance, "balance.txt")
		default:
			fmt.Print("Thankyou for visiting Go Bank !")
			return
		}
	}
}
