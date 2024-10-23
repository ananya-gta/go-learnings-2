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
	title     string
	content   string
	createdAt time.Time
}

// constructor function
func New(title, content string) (NoteStruct, error) {
	if title == "" || content == "" {
		return NoteStruct{}, errors.New("invalid input")
	}
	return NoteStruct{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (note NoteStruct) Save() (error){
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName)
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	os.WriteFile(fileName, json, 0644)

}

func (note NoteStruct) DisplayNote() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v", note.title, note.content)
}
