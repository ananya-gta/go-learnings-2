package main

import (
	"fmt"

	// "example.com/price-calculator/cmdManager"
	"example.com/price-calculator/fileManager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.01, 0.07, 0.11}

	for _, taxRate := range taxRates {
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmd := cmdManager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("couple not process job:", err)
		}
	}

	// now you can swap between fm and cmd

}


// func main() {
// 	prices := []float64{10, 20, 30}
// 	taxRates := []float64{0, 0.7, 0.1, 0.15}

// 	result := make(map[float64][]float64)

// 	for _, taxRate := range taxRates {

// 		taxIncludedPrices := make([]float64, len(prices))
// 		for priceIndex, price := range prices {
// 			taxIncludedPrices[priceIndex] = price * (1 + taxRate)
// 		}

// 		result[taxRate] = taxIncludedPrices
// 	}

// 	fmt.Println(result)
// }