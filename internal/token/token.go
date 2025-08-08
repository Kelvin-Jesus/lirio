package token

import "fmt"

type TokenType string

// Constants for token types
const (
	// Single-char tokens
	TOK_LEFT_PAREN     TokenType = "TOK_LEFT_PAREN"     // (
	TOK_RIGHT_PAREN    TokenType = "TOK_RIGHT_PAREN"    // )
	TOK_LEFT_CURLY     TokenType = "TOK_LEFT_CURLY"     // {
	TOK_RIGTH_CURLY    TokenType = "TOK_RIGTH_CURLY"    // }
	TOK_PLUS_SIGN      TokenType = "TOK_PLUS_SIGN"      // +
	TOK_MINUS_SIGN     TokenType = "TOK_MINUS_SIGN"     // -
	TOK_STAR_SIGN      TokenType = "TOK_STAR_SIGN"      // *
	TOK_LEFT_SQR_BCKT  TokenType = "TOK_LEFT_SQR_BCKT"  // [
	TOK_RIGHT_SQR_BCKT TokenType = "TOK_RIGHT_SQR_BCKT" // ]
	TOK_DOT            TokenType = "TOK_DOT"            // .
	TOK_COMMA          TokenType = "TOK_COMMA"          // ,
	TOK_CARET          TokenType = "TOK_CARET"          // ^
	TOK_SLASH          TokenType = "TOK_SLASH"          // /
	TOK_SEMICOLON      TokenType = "TOK_SEMICOLON"      // ;
	TOK_QUESTION       TokenType = "TOK_QUESTION"       // ?
	TOK_MOD_SIGN       TokenType = "TOK_MOD_SIGN"       // %
)

type Token struct {
	Type   TokenType
	Lexeme []rune
	// The line number of the token
	Line uint
}

func (token *Token) AsString() string {
	return fmt.Sprintf("(Lexeme: %c, Type: %s, Line: %d)", token.Lexeme, token.Type, token.Line)
}
