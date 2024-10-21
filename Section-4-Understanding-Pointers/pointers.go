package main

import "fmt"

func main() {
	age := 32
	// var agePointer *int
	agePointer := &age
	agePointerValue := *agePointer
	
	fmt.Println(agePointer)
	fmt.Println(agePointerValue)
	
	fmt.Println(getAdultYears(age, agePointer))
	fmt.Println(age)
}

func getAdultYears(age int, agePointer *int) (int, int) {
	return age, *agePointer - 18
}