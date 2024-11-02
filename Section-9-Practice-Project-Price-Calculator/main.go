package main

// import "fmt"
import ("example.com/price-calculator/prices")

func main() {
	taxRates := []float64{0, 0.1, 0.7, 0.11}
	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
	
}
