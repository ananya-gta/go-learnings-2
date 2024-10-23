package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note-taking-app/noteStruct"
	"example.com/note-taking-app/todo"
)

func main() {
	title, content := getNoteData() 
	userNote, err := noteStruct.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.DisplayNote()

	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving the note succeeded")

	todoText := getTodoData()

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	todo.DisplayTodo()

	err = todo.Save()
	if err != nil {
		fmt.Println("Saving the todo failed")
		return
	}
	fmt.Println("Saving the todo succeeded")
		
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

func getTodoData() (string) {
	text := getUserInput("Todo Text: ")
	return text
}


