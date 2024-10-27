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

	//here you will not be passing a named func but an anonymous function as argument
	// defining function just in time
	doubleNumbersAnonymously(&numbers, func(number int) int { return number * 2 })

	// createTransformer is now a factory function which produces functions
	// you can create a double function like this
	double1 :=  factoryFunction(2)
	fmt.Println(double1(10))
	// or triple function like this
	triple1 := factoryFunction(3)
	fmt.Println(triple1(15))
	doubledThroughFactoryFunction := transformNumbers(&numbers, double1)
	fmt.Println(doubledThroughFactoryFunction)
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

// anonymous functions
func doubleNumbersAnonymously(numbers *[]int, transform func(int) int) []int {
	doubleNumbers := []int{}
	for _, number := range *numbers {
		doubleNumbers = append(doubleNumbers, transform(number))
	}
	return doubleNumbers
}

// the below function produces a function
func factoryFunction(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}


