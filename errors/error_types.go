package errors

import (
	"fmt"
)

type FileNotFoundError struct {
	File string
	Err  error
}

func (m FileNotFoundError) Error() string {
	return fmt.Sprintf("map: %v", m.File+m.Err.Error())
}

func (m FileNotFoundError) getFile() string {
	return m.File
}

func MapFileNotFoundError(file string, err error) *FileNotFoundError {
	return &FileNotFoundError{file, err}
}
