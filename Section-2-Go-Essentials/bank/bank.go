package main

import "fmt"

func main() {
	printChoices()
}

func printChoices() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Enter your choice from the below options.")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}
