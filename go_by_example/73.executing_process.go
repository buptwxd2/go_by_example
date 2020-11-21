package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main()  {
	binary, lookError := exec.LookPath("ls")
	if lookError != nil{
		panic(lookError)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
