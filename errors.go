package goenvloader

import "fmt"

type illegalVariableDeclared struct {
	token string
}

func (e *illegalVariableDeclared) Error() string {
	return fmt.Sprintf("Illegal Variable declared. Assignment not found in %s", e.token)
}

type emptyToken struct{}

func (e *emptyToken) Error() string {
	return "Empty variable supplied"
}
