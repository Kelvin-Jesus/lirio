package token

import "fmt"

type TokenType string

// Constants for token types
const (
	// Single-char tokens
	TOK_LEFT_PAREN  TokenType = "TOK_LEFT_PAREN"  // (
	TOK_RIGHT_PAREN TokenType = "TOK_RIGHT_PAREN" // )
	TOK_LEFT_CURLY  TokenType = "TOK_LEFT_CURLY"  // {
	TOK_RIGTH_CURLY TokenType = "TOK_RIGTH_CURLY" // }
	TOK_PLUS_SIGN   TokenType = "TOK_PLUS_SIGN"   // +
	TOK_MINUS_SIGN  TokenType = "TOK_MINUS_SIGN"  // -
	TOK_STAR_SIGN   TokenType = "TOK_STAR_SIGN"   // *

)

type Token struct {
	Type   TokenType
	Lexeme []rune
}

func (token *Token) AsString() string {
	return fmt.Sprintf("(Lexeme: %c, Type: %s)", token.Lexeme, token.Type)
}
