package main

import "fmt"

func main() {
    fmt.Println("Q1_Result: ", q1())
    fmt.Println("Q2_Result: ", q2())
    fmt.Println("Q3_Result: ", q3())
    fmt.Println("Q4_Result: ", q4())
    fmt.Print("Q5_Result, After swapping: ")
	a, b := q5()
	fmt.Println(a, b)
	fmt.Println("Q6_Result: ", q6())
	fmt.Println("Q7_Result: ", q7())
	fmt.Print("Q8_Result, After updating: ")
	c, d := q8()
	fmt.Println(c, d)
}
