package main

import "fmt"
import "errors"
import "example.com/note-taking-app/noteStruct"

func main() {
	title, content, err := getNoteData() 
	if err != nil {
		fmt.Println(err)
		return
	}
	
}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)
	if value == "" {
		return "", errors.New("invalid")
	}
	return value, nil
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note Title: ")

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	content, err := getUserInput("Note content: ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return title, content, nil
}
