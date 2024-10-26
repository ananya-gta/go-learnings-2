package main
import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	fmt.Println(productNames)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	fmt.Println(prices[1:3])
	fmt.Println(prices[:3])

	featuredPrices := prices[1:]
	fmt.Println(featuredPrices)
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)

	fmt.Println(len(featuredPrices), cap(featuredPrices))
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
	dynamicArray()
	hobbies()
}

func dynamicArray() {
	dynamic := []float64{1, 2}
	dynamic[1] = 19.99
	fmt.Println(append(dynamic, 2.99))
	fmt.Println(dynamic)
}

func hobbies() {
	hobbies :=  []string{"reading", "swimming", "gaming"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])
	slice1 :=  hobbies[0:2]
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	slice1 = append(slice1, hobbies[2])
	slice1 = slice1[1:]
	fmt.Println(slice1)

	
	


}