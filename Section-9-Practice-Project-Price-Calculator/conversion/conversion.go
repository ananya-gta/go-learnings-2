package conversion

import (
	"errors"
	"strconv"
)

func ConvertStringToFloat(stringSlice []string) ([]float64, error) {
	// Convert the input string to a float64
	var floatSlice []float64
	for _, stringVal := range stringSlice {
		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("failed to convert string into float")
		}

		floatSlice = append(floatSlice, floatVal)

	}
	return floatSlice, nil
}
