package main

import (
	"fmt"
	"io"
	"os"
	"rug/internal/lexer"
)

func main() {

	// Check if the filename is provided
	if len(os.Args) < 2 {
		fmt.Println("./rug <filename>")
		return
	}

	// Open the file
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file content
	src, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Tokenize the source code
	toks := lexer.Tokenize(string(src))

	for token := range toks {
		fmt.Println(token)
	}
}
