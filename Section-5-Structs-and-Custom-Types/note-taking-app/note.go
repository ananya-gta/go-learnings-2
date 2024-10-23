package main

import "fmt"
import "errors"

func main() {
	getUserInput()

}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)
	if (value == "") {
		return "", errors.New("invalid")
	}
	return value, nil
}