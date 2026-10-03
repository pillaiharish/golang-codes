// "Write a function that reads network configuration and distinguishes missing file from other errors."

package main

import (
	"errors"
	"fmt"
	"os"
)

func loadFile(filePath string) ([]byte, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("\n\t\"Error:\" { \"Error while file %w \"}", err)
	}
	return fileData, nil
}

func main() {
	_, err := loadFile("./config.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File Not Found:", err)
		} else {
			fmt.Println("Unexpected Failure", err)
		}
	}

}

// os.ReadFile()
//      │
//      │ returns ErrNotExist
//      ▼
// loadConfig()
//      │
//      │ wraps it with context
//      ▼
// "load config ./config.json: file does not exist"
//      │
//      │ errors.Is traverses chain
//      ▼
// os.ErrNotExist found

// Output:
// File Not Found:
//         "Error:" { "Error while file open ./config.json: no such file or directory "}
