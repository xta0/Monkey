package parser

import (
	"go/token"

	"../ast"
	"../lexer"
	"../token"
)

// Parser Definination
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New a parser by passing in a lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//Read tokens
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parser's main entrance
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetToken()
	default:
		return nil
	}
}
func (p *Parser) parseLetToken() *ast.LetStatement {
	stmt := &ast.LetStatement{}
	stmt.Token = p.curToken
	if !p.expectedPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectedPeek(token.ASSIGN) {
		return nil
	}
	//TODO: stmt.Value
	if !p.currentTokenIs(token.SEMICOLON) {
		return nil
	}
	p.nextToken()
	return stmt
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectedPeek(t token.TokenType) bool {
	if !p.peekTokenIs(t) {
		return false
	}
	p.nextToken()
	return true
}
