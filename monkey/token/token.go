package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

//token types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier + literals
	IDENT = "IDENT" // add,foobar,x,y,...
	INT   = "INT"   //12345

	//Operators
	ASSIGN = "="
	PLUS   = "+"

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
)
