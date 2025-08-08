package lexer

import (
	"unicode"

	"github.com/Kelvin-Jesus/lirio/internal/token"
)

// start and current keeps track of the
// range of characters beeing tokenized
var (
	start   int = 0
	current int = 0
)

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
		case '=':
			// is "=="
			if lexer.match('=') {
				lexer.addToken(token.TOK_EQ)
			} else {
				lexer.addToken(token.TOK_ASSIGN_OP)
			}
		case '!':
			if lexer.match('=') {
				lexer.addToken(token.TOK_NEQ)
			} else {
				lexer.addToken(token.TOK_NOT)
			}
		case '<':
			if lexer.match('=') {
				lexer.addToken(token.TOK_LESS_OR_EQ)
			} else {
				lexer.addToken(token.TOK_LESS)
			}
		case '>':
			if lexer.match('=') {
				lexer.addToken(token.TOK_GREATER_OR_EQ)
			} else {
				lexer.addToken(token.TOK_GREATER)
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			lexer.handleNumber()
		}
		// @TODO If is the ' token, apply logic to read an entire string
		// @TODO iF is an alpha caracther, _ or ?, then treat as a identifier

	}
}

func (lexer *Lexer) handleNumber() {
	for unicode.IsDigit(lexer.peek()) {
		lexer.advance()
	}

	// if the next token is . and the next is another digit, if it is
	// that's a float
	if lexer.peek() == '.' && unicode.IsDigit(lexer.lookAhead(1)) {
		lexer.advance() // consumes the '.'
		for unicode.IsDigit(lexer.peek()) {
			lexer.advance()
		}

		lexer.addToken(token.TOK_FLOAT)
	} else {
		lexer.addToken(token.TOK_INTEGER)
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
