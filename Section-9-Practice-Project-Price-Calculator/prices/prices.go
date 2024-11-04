package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/fileManager"
)

type TaxIncludedPriceJob struct {
	IOManager         fileManager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func NewTaxIncludedPriceJob(fm fileManager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: fm,
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

	job.TaxIncludedPrices = result

	job.IOManager.WriteFile(job.TaxIncludedPrices)

}

func (job *TaxIncludedPriceJob) LoadData() error {
	//read lines from file
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println("error occured while reading lines from file: ", err)
		return err
	}

	// convert each string line into float64 and store it into InputPrices
	prices, err := conversion.ConvertStringToFloat(lines)

	if err != nil {
		fmt.Println("error occured while converting string to float64: ", err)
		return err
	}

	job.InputPrices = prices

	return nil

}
