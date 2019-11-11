package ast

import "../token"

// Node xxx
type Node interface {
	TokenLiteral() string
}

// Statement xxx
type Statement interface {
	Node
	statementNode()
}

// Expression xxx
type Expression interface {
	Node
	expressionNode()
}

// Identifier implements the EXpression interface
type Identifier struct {
	Token token.Token
	Value string
}

// Program is the root node of our AST
type Program struct {
	Statements []Statement
}

// TokenLiteral xxx
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement xxx
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// TokenLiteral - Interface implementation
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) statementNode() {}
