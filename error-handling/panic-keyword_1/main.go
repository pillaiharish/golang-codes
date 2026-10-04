package main 

import (
	"fmt"
	"os"
)

func fileReader() int{
	file, err := os.Open("non_existing_file.json")
	if err != nil {
		panic("There is no need to move forward as the file open failed.")
	}
	defer file.Close()
	fmt.Println("File opened successfully")
	return 0
}

func main() {
	response := fileReader()
	fmt.Println("Response is:", response)
}

// Output:
// panic: There is no need to move forward as the file open failed.
//
// goroutine 1 [running]:
// main.fileReader()
//   ~/golang-codes/error-handling/panic-keyword_1/main.go:11 +0xcc
// main.main()
// 	 ~/golang-codes/error-handling/panic-keyword_1/main.go:19 +0x1c
// exit status 2
