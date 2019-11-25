package parser

import (
	"go/ast"

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
func (p *Parser) parseLetToken() ast.Statement {
	stmt := ast.LetStatement{}
	stmt.Token = p.curToken
	return nil
}
