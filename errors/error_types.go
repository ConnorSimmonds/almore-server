package errors

import (
	"errors"
	"fmt"
)

type MapError struct {
	err  error
	code uint8
}

func (m MapError) Error() string {
	return fmt.Sprintf("map: %v", m.err)
}

func ReturnMapFileError() *MapError {
	return &MapError{
		errors.New("FileNotFound"),
		0,
	}
}

func ReturnMapNil() *MapError {
	return nil
}
