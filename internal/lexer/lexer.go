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
var currentLine uint = 1

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

		switch character {
		case '\n':
			currentLine++
		case '\t', '\r':
			continue
		case '/':
			// it's a comment
			if lexer.peek() == '/' {
				// consumes tokens until it reaches \n
				for lexer.peek() != '\n' {
					lexer.advance()
				}
			} else {
				lexer.addToken(token.TOK_SLASH)
			}
		case '(':
			lexer.addToken(token.TOK_LEFT_PAREN)
		case ')':
			lexer.addToken(token.TOK_RIGHT_PAREN)
		case '[':
			lexer.addToken(token.TOK_LEFT_SQR_BCKT)
		case ']':
			lexer.addToken(token.TOK_RIGHT_SQR_BCKT)
		case '.':
			lexer.addToken(token.TOK_DOT)
		case ',':
			lexer.addToken(token.TOK_COMMA)
		case '+':
			lexer.addToken(token.TOK_PLUS_SIGN)
		case '-':
			lexer.addToken(token.TOK_MINUS_SIGN)
		case '*':
			lexer.addToken(token.TOK_STAR_SIGN)
		case '^':
			lexer.addToken(token.TOK_CARET)
		case ';':
			lexer.addToken(token.TOK_SEMICOLON)
		case '?':
			lexer.addToken(token.TOK_QUESTION)
		case '%':
			lexer.addToken(token.TOK_MOD_SIGN)
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
			Line:   currentLine,
		},
	)
}

// Takes a peek at the current character but dsn't consume it
func (lexer *Lexer) peek() rune {
	return lexer.Source[current]
}

// Takes a peek the character at current + *position*
// without consuming it
func (lexer *Lexer) lookAhead(position int) rune {
	return lexer.Source[current+position]
}

// Check if the character at current position matches
// the *expected* and if so, consumes the character
func (lexer *Lexer) match(expected rune) bool {
	if current >= len(lexer.Source) {
		return false
	}

	if lexer.Source[current] != expected {
		return false
	}

	current++

	return true
}
