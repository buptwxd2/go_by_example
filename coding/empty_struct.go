package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type myStruct struct {
	bb string
}

func main()  {

	a := myStruct{bb:"a"}

	s := struct {
	}{}

	fmt.Println("Type of s:", reflect.TypeOf(a))
	fmt.Println("Size of s:", unsafe.Sizeof(s))
}