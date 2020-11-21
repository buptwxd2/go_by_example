package main

import "os"

func main()  {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != err{
		panic(err)
	}
}