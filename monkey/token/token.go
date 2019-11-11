package token

// TokenType xxx
type TokenType string

// Token xxx
type Token struct {
	Type    TokenType
	Literal string
}

// supported token types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier + literals
	IDENT = "IDENT" // add,foobar,x,y,...
	INT   = "INT"   //12345

	//Operators
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	BANG     = "!"
	ASSIGN   = "="
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

// LookupIndent returns the token type of indent
func LookupIndent(ident string) TokenType {
	if identType, ok := keywords[ident]; ok {
		return identType
	}
	return IDENT
}
