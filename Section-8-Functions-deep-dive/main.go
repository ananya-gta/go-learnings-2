package main

import "fmt"

// custom type
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := transformNumbers(&numbers, double) 
	tripled := transformNumbers(&numbers, triple) 
	fmt.Println("After doubling: \n", doubled)             
	fmt.Println("After tripling: \n", tripled)             
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	transformNumbers := []int{}
	for _, number := range *numbers {
		transformNumbers = append(transformNumbers, transform(number))
	}
	return transformNumbers 
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}







