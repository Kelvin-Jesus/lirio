package main

import (
	"fmt"
	"os"

	"github.com/Kelvin-Jesus/lirio/internal/lexer"
	"github.com/Kelvin-Jesus/lirio/internal/token"
)

func main() {
	if len(os.Args) != 2 {
		panic("Usage: go run/build main.go <filename>")
	}

	filename := os.Args[1]
	fmt.Println(filename)

	sourceCode, err := ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nLexer:")

	var tokens []token.Token
	lxr := lexer.Lexer{
		Source: []rune(string(sourceCode)),
		Tokens: tokens,
	}

	lxr.Tokenize()

	for _, tok := range lxr.Tokens {
		fmt.Println(tok.AsString())
	}
}

func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
