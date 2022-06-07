package goenvloader

import (
	"bufio"
	"os"
	"path/filepath"
)

// loadEnvFile returns an io.ReadCloser object curresponding to the env file.
// It returns an error, if the specified file cannot be opened due to some reason.
// The function builds the path of the file with the given name and the working directory
func loadEnvFile(filename string) (*bufio.Scanner, error) {
	pwd, _ := os.Getwd()
	var path string = filepath.Join(pwd, filename)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	} else {
		scanner := createNewLineScanner(file)
		return scanner, nil
	}
}
