package main

import "fmt"
import "example.com/note-taking-app/noteStruct"
import "bufio"
import "os"
import "strings"

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
	// var value string
	// fmt.Scanln(&value) 
	// the above code will display only the first word, it cannot display long sapce separated text
	reader := bufio.NewReader(os.Stdin)

	text, err :=reader.ReadString('\n')
	// reader.ReadString('\n') will return a string with a newline character at the end

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content:= getUserInput("Note content: ")
	return title, content
}


