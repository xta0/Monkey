package lexer

import (
	"../token"
)

// Lexer Definition
type Lexer struct {
	input        string
	position     int  //current position in input
	readPosition int  //current reading position in input
	ch           byte //current char under examination
}

// New function create a new lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readchar()
	return l
}

// NextToken returns the next Lexer token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekchar() == '=' {
			ch := l.ch
			l.readchar()
			tok = newToken(token.EQ, string(ch)+string(l.ch))
		} else {
			tok = newToken(token.ASSIGN, string(l.ch))
		}
	case ',':
		tok = newToken(token.COMMA, string(l.ch))
	case '+':
		tok = newToken(token.PLUS, string(l.ch))
	case '-':
		tok = newToken(token.MINUS, string(l.ch))
	case '*':
		tok = newToken(token.ASTERISK, string(l.ch))
	case '/':
		tok = newToken(token.SLASH, string(l.ch))
	case '!':
		if l.peekchar() == '=' {
			ch := l.ch
			l.readchar()
			tok = newToken(token.NOT_EQ, string(ch)+string(l.ch))
		} else {
			tok = newToken(token.BANG, string(l.ch))
		}
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch))
	case ')':
		tok = newToken(token.RPAREN, string(l.ch))
	case '(':
		tok = newToken(token.LPAREN, string(l.ch))
	case '{':
		tok = newToken(token.LBRACE, string(l.ch))
	case '}':
		tok = newToken(token.RBRACE, string(l.ch))
	case '>':
		tok = newToken(token.GT, string(l.ch))
	case '<':
		tok = newToken(token.LT, string(l.ch))
	case 0:
		tok = newToken(token.EOF, "")
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIndent(tok.Literal)
			return tok
		} else if isNumber(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, string(l.ch))
		}
	}
	l.readchar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readchar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isNumber(l.ch) {
		l.readchar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readchar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
		return
	}
	l.ch = l.input[l.readPosition]
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekchar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readchar()
	}
}

func newToken(tokenType token.TokenType, litral string) token.Token {
	return token.Token{Type: tokenType, Literal: litral}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch >= 'Z' || ch == '_'
}
func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
