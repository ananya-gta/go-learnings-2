package fileManager

// this file will read the data from a file
import (
	"bufio"
	"errors"
	"os"
	"encoding/json"
)

type FileManager struct {
	InputFilepath string
	OutputFilepath string
}

// below it does not matter you use a pointer as a receiver or simple struct because here we are just reading the data not updating any value
func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilepath)
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


func (fm FileManager) WriteFile(data interface{}) error {
	file, err := os.Create(fm.OutputFilepath)

	if err != nil {
		return errors.New("failed to create a file")
	}

	defer file.Close() // Ensure the file is closed after we're done

    encoder := json.NewEncoder(file)
    if err := encoder.Encode(data); err != nil {
        return errors.New("failed to convert data into JSON: " + err.Error())
    }

    return nil
}

func New(inputPath, outputPath string) FileManager{
	return FileManager{
		InputFilepath: inputPath,
		OutputFilepath: outputPath,
	}
}
