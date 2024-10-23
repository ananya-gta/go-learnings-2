package noteStruct

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"encoding/json"
)

type NoteStruct struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

// constructor function
func New(title, content string) (NoteStruct, error) {
	if title == "" || content == "" {
		return NoteStruct{}, errors.New("invalid input")
	}
	return NoteStruct{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note NoteStruct) Save() (error){
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
	
}

func (note NoteStruct) DisplayNote() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v", note.Title, note.Content)
}
