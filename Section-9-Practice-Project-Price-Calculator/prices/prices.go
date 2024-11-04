package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	// "example.com/price-calculator/fileManager"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	/* The backtick notation (`json:"-"`) is a struct tag.
	In this case, json:"-" tells the JSON marshalling/unmarshalling process to ignore this field
	when converting the struct to or from JSON.
	This means that when an instance of TaxIncludedPriceJob is encoded to JSON,
	the IOManager field will not be included in the resulting JSON object.
	Similarly, when decoding JSON into the struct, any JSON properties corresponding to IOManager will be ignored. */
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) Process() error{
	// read prices from txt file
	err := job.LoadData()
	if err != nil {
		return  err

	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	return job.IOManager.WriteFile(job)

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
