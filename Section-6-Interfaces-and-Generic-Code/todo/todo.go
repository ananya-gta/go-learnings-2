package todo

import (
	"errors"
	"fmt"
	"os"
	"encoding/json"
)

type Todo struct {
	
	Text   string `json:"text"`
}

// constructor function
func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Save() (error){
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
	
}

func (todo Todo) DisplayTodo() {
	fmt.Printf(todo.Text)
}
