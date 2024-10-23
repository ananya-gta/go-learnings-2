package main

import "fmt"
import "example.com/note-taking-app/noteStruct"

func main() {
	title, content := getNoteData() 
	userNote, err := noteStruct.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.DisplayNote()
	
}

func getUserInput(prompt string) (string) {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)
	return value
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content:= getUserInput("Note content: ")
	return title, content
}


