package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	// Open function is used to open a file and load it
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open file.")
	}

	// A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	defer file.Close()

	// bufio.NewScanner is used to read line by line from an open file
	scanner := bufio.NewScanner(file)

	var lines []string

	// scan returns a boolean i.e. false if there is nothing to scan
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// this will simply return an error, if something went wrong
	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, errors.New("Failed to read line in file.")
	}

	// file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	// Create function is used to create a file
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	defer file.Close()

	// this is added to simulation a slow running process
	time.Sleep(3 * time.Second)

	// this is to write data in a file, pass the file here
	encoder := json.NewEncoder(file)
	// pass the data here
	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New("Faild to convert data to JSON.")
	}

	// file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
