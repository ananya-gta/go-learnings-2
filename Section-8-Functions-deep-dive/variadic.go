package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	sumUp := sum(numbers)
	fmt.Println(sumUp)
	fmt.Println(variadicSum(1,2,3,4,4,5,5))
	// Combine 1 with the numbers slice
	combined := append([]int{1}, numbers...) // Create a new slice
	anotherSum := sum(combined)
	fmt.Println(anotherSum)

	anotherSum1 := variadicSum(numbers...)
	fmt.Println(anotherSum1)
}

func sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// you can work with functions that have any amount of parameters
func variadicSum(numbers ...int) int {
	// these three dots, collect all the standalone values and merge all of them into a slice
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

/* from above, numbers is a collect all parameter
reverse of this will be 
numbers... --> this will pull all the values from the slice numbers 
and turn slice into a list of standalone values */
