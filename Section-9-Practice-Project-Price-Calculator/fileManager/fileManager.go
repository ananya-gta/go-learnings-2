package fileManager

// this file will read the data from a file
import (
	"bufio"
	"errors"
	"os"
	"encoding/json"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {

		return nil, errors.New("an eror occured while opening the file")
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("an error occured while reading the file")
	}

	file.Close()
	return lines, nil
}


func WriteFile(path string,  data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create a file")
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to convert data into JSON")
	}

	file.Close()
	return nil
}

