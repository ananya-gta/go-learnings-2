package main

import (
	"fmt"

	"example.com/price-calculator/cmdManager"
	"example.com/price-calculator/fileManager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.01, 0.07, 0.11}
	
	for _, taxRate := range taxRates {
		fm :=  fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmd := cmdManager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
	
}
