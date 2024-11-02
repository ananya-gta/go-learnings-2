package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/fileManager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) Process() {
	// read prices from txt file
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)
}

func (job TaxIncludedPriceJob) LoadData() {
	//read lines from file
	lines, err := fileManager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println("error occured while reading lines from file: ", err)
		return
	}

	// convert each string line into float64 and store it into InputPrices
	prices, err := conversion.ConvertStringToFloat(lines)

	if err != nil {
		fmt.Println("error occured while converting string to float64: ", err)
		return
	}

	job.InputPrices = prices

}
