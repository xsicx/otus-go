package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatal("Not enough arguments. Usecase: go-envdir /path/to/env/dir command arg1 arg2")
	}

	dir, cmd := args[0], args[1:]

	envs, err := ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(RunCmd(cmd, envs))
}
