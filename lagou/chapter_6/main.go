package main

import (
	"fmt"
	"math"
)

//factory function
func New(text string) *errorString {
	return &errorString{text}
}


type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func main()  {
	// type assertion
	var s fmt.Stringer
	s =
}
