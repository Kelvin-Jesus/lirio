package lexer

import (
	"github.com/Kelvin-Jesus/lirio/internal/token"
)

// start and current keeps track of the
// range of characters beeing tokenized
var start int = 0
var current int = 0

// Keeps tracking of the current line
// of source code
var currentLine int = 1

type Lexer struct {
	Tokens []token.Token
	Source []rune
}

func (lexer *Lexer) Tokenize() {
	var sourceLength int = len(lexer.Source)
	for current < sourceLength {
		// they start at the same character and
		// current increments until a valid token
		// is found
		start = current

		var character rune = lexer.advance()

		if character == '+' {
			lexer.addToken(token.TOK_PLUS_SIGN)
		}

		if character == '-' {
			lexer.addToken(token.TOK_MINUS_SIGN)
		}

		if character == '*' {
			lexer.addToken(token.TOK_STAR_SIGN)
		}

	}
}

// Advances the current pointer consuming a character
func (lexer *Lexer) advance() rune {
	currentCharacter := lexer.Source[current]

	current++

	return currentCharacter
}

// Add a new token of *tokenType* to the list of tokens
func (lexer *Lexer) addToken(tokenType token.TokenType) {
	lexer.Tokens = append(
		lexer.Tokens,
		token.Token{
			Type:   tokenType,
			Lexeme: lexer.Source[start:current],
		},
	)
}
