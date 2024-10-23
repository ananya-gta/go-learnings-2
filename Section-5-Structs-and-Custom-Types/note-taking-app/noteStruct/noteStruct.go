package noteStruct

import (
	"errors"
	"fmt"
	"time"
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
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}

func (note NoteStruct) DisplayNote() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v", note.title, note.content)
}
