package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteFile(data any) error
}