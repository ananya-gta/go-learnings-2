package main

import "fmt"

// custom type
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{7, 2, 3, 4, 5}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println("After doubling: \n", doubled)
	fmt.Println("After tripling: \n", tripled)

	//calling the function which returns a function
	transformFun1 := getTransformedFunc(&numbers)
	transformFun2 := getTransformedFunc(&moreNumbers)

	//gettting the new array by passing an array and function to the function
	tranformNumbers := transformNumbers(&numbers, transformFun1)
	tranformMoreNumbers := transformNumbers(&moreNumbers, transformFun2)
	
	fmt.Println(tranformNumbers)
	fmt.Println(tranformMoreNumbers)
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

// when you execute the below function you get a func as returned value
func getTransformedFunc(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}
