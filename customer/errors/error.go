package errors

import "fmt"

type Error struct {
	Description string `json:"description"`
	StatusCode  int    `json:"status_code"`
}

func Define(desc string, statusCode int) Error {
	return Error{
		Description: desc,
		StatusCode:  statusCode,
	}
}

func (e *Error) Log() {
	fmt.Printf(e.Description+", status-code: %d", e.StatusCode)
}
