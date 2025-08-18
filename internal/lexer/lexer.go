package lexer

import (
	"fmt"
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

		if character == '\n' {
			currentLine++
		} else if character == '\t' || character == '\r' {
			continue
		} else if character == '/' {
			// it's a comment
			if lexer.peek() == '/' {
				// consumes tokens until it reaches \n
				for lexer.peek() != '\n' {
					lexer.advance()
				}
			} else {
				lexer.addToken(token.TOK_SLASH)
			}
		} else if character == '(' {
			lexer.addToken(token.TOK_LEFT_PAREN)
		} else if character == ')' {
			lexer.addToken(token.TOK_RIGHT_PAREN)
		} else if character == '[' {
			lexer.addToken(token.TOK_LEFT_SQR_BCKT)
		} else if character == ']' {
			lexer.addToken(token.TOK_RIGHT_SQR_BCKT)
		} else if character == '.' {
			lexer.addToken(token.TOK_DOT)
		} else if character == ',' {
			lexer.addToken(token.TOK_COMMA)
		} else if character == '+' {
			lexer.addToken(token.TOK_PLUS_SIGN)
		} else if character == '-' {
			lexer.addToken(token.TOK_MINUS_SIGN)
		} else if character == '*' {
			lexer.addToken(token.TOK_STAR_SIGN)
		} else if character == '^' {
			lexer.addToken(token.TOK_CARET)
		} else if character == ';' {
			lexer.addToken(token.TOK_SEMICOLON)
		} else if character == '?' {
			lexer.addToken(token.TOK_QUESTION)
		} else if character == '%' {
			lexer.addToken(token.TOK_MOD_SIGN)
		} else if character == '=' {
			// is "=="
			if lexer.match('=') {
				lexer.addToken(token.TOK_EQ)
			} else {
				lexer.addToken(token.TOK_ASSIGN_OP)
			}
		} else if character == '!' {
			if lexer.match('=') {
				lexer.addToken(token.TOK_NEQ)
			} else {
				lexer.addToken(token.TOK_NOT)
			}
		} else if character == '<' {
			if lexer.match('=') {
				lexer.addToken(token.TOK_LESS_OR_EQ)
			} else {
				lexer.addToken(token.TOK_LESS)
			}
		} else if character == '>' {
			if lexer.match('=') {
				lexer.addToken(token.TOK_GREATER_OR_EQ)
			} else {
				lexer.addToken(token.TOK_GREATER)
			}
		} else if unicode.IsDigit(character) {
			lexer.handleNumber()
		} else if character == '\'' {
			//single quote "'"
			lexer.handleString()
		} else if character == 'l' {
			lexer.handleIdentifier()
		}
	}
}

func (lexer *Lexer) handleIdentifier() {
	if lexer.peek() == 'e' && lexer.lookAhead(1) == 't' {
		lexer.advance()
		lexer.advance()
		lexer.advance()
	}

	// if starts with ?, it's not valid
	if lexer.peek() == '?' {
		panic(
			fmt.Sprintf(
				"error in line: [%d] -> A Identifier should not start with the '?' character",
				currentLine,
			),
		)
	}

	for unicode.IsLetter(lexer.peek()) || unicode.IsDigit(lexer.peek()) || lexer.peek() == '_' || lexer.peek() == '?' && lexer.peek() != '\n' {
		lexer.advance()
	}

	// check if identifier matches any keyword from hashmap
	// sum start with 4 'cause of the identifier is
	// 'let '
	isLetIdentifier := string(lexer.Source[start:current]) == "let"

	currentText := lexer.Source[start+4 : current]
	fmt.Println(string(currentText))
	if _, ok := token.Keywords[string(currentText)]; ok {
		if isLetIdentifier {
			panic(
				fmt.Sprintf(
					"error in line [%d] -> '%s' is a keyword and should not be used as an identifier",
					currentLine,
					string(currentText),
				),
			)
		}
	}

	lexer.addToken(token.TOK_IDENTIFIER)
}

func (lexer *Lexer) handleString() {
	for lexer.peek() != '\'' {
		lexer.advance()
	}

	lexer.advance()
	lexer.addToken(token.TOK_STRING)
}

func (lexer *Lexer) handleNumber() {
	for unicode.IsDigit(lexer.peek()) {
		lexer.advance()
	}

	// if the next token is . and the next is another digit,
	// if true, that's a float
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
