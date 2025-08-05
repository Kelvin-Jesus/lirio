package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("Usage: go run/build main.go <filename>")
	}

	filename := os.Args[1]
	fmt.Println(filename)

	ReadFile(filename)

	// @TODO Tokenize source code
}

func ReadFile(filePath string) {
	// sourceCodeFileDescriptor, error := os.Open(filePath)
	// if error != nil {
	// 	panic("error opening source file")
	// }

	file, _ := os.ReadFile(filePath)
	fmt.Println(string(file))
}
