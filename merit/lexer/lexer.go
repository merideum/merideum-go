package lexer

import (
	"github.com/merideum/merideum-go/merit/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	nextPosition int  // current reading position in input (after current char)
	currentChar  byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.nextPosition]
	}
	l.position = l.nextPosition

	l.nextPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.currentChar {
	case '=':
		tok = newToken(token.ASSIGN, l.currentChar)
	// case ';':
	// 	tok = newToken(token.SEMICOLON, l.ch)
	// case '(':
	// 	tok = newToken(token.LPAREN, l.ch)
	// case ')':
	// 	tok = newToken(token.RPAREN, l.ch)
	// case ',':
	// 	tok = newToken(token.COMMA, l.ch)
	// case '+':
	// 	tok = newToken(token.PLUS, l.ch)
	// case '{':
	// 	14
	// 	tok = newToken(token.LBRACE, l.ch)
	// case '}':
	// 	tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.currentChar) {
			tok.Type = token.INTEGER
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.currentChar)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}
