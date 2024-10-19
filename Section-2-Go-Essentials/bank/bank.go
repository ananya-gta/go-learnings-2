package main

import "fmt"

func main() {
	var accountBalance = 1000.0
	fmt.Println("Welcome to Go Bank!")
	startBankApplication(accountBalance)
}

func printChoices() {
	fmt.Println("Enter your choice from the below options.")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
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
			default:
				fmt.Print("Thankyou for visiting Go Bank !")
				return
		}
	}
}
