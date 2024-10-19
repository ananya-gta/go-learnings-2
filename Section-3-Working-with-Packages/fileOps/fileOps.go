package fileOps

import "fmt"
import "os"
import "errors"
import "strconv"

func ReadFloatValueFromFile(filename string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}
	value, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return defaultValue, errors.New("failed to parse stored value")
	}
	return value, nil
}

func WriteIntoFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}