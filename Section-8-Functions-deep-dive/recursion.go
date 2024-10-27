package main

import "fmt"

func recursion() {
	fmt.Println(factorial(5))
}

func factorial(number int) int{
	if number ==  0 {
		return 1
	}

	return  number * factorial(number - 1)

}