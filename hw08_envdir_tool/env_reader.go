package main

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrIncorrectDirectory      = errors.New("incorrect directory")
	ErrIncorrectEnvDirFilename = errors.New("incorrect filename in env dir")
	ErrIncorrectEnvDirFile     = errors.New("incorrect file in env dir")
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, ErrIncorrectDirectory
	}

	envMap := make(Environment)

	for _, file := range files {
		err = validateFile(file)
		if err != nil {
			// log error if we need
			continue
		}

		envValue, err := getEnvValueFromFile(dir, file.Name())
		if err != nil {
			// log error if we need
			continue
		}

		envMap[file.Name()] = envValue
	}

	return envMap, nil
}

func getEnvValueFromFile(dir, fileName string) (EnvValue, error) {
	envValue := EnvValue{}
	file, err := os.Open(filepath.Join(dir, fileName))
	if err != nil {
		return envValue, ErrIncorrectEnvDirFile
	}
	defer file.Close()
	r := bufio.NewReader(file)

	fileInfo, err := file.Stat()
	if err != nil {
		return EnvValue{}, err
	}

	if fileInfo.Size() == 0 {
		envValue.NeedRemove = true
		return envValue, nil
	}

	value, _, err := r.ReadLine()
	if err != nil {
		return EnvValue{}, err
	}

	value = bytes.ReplaceAll(value, []byte("\x00"), []byte("\n"))
	envValue.Value = string(bytes.TrimRight(value, " "))

	return envValue, nil
}

func validateFile(file os.DirEntry) error {
	if strings.Contains(file.Name(), "=") {
		return ErrIncorrectEnvDirFilename
	}

	if !file.Type().IsRegular() {
		return ErrIncorrectEnvDirFile
	}

	return nil
}
