package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	fmt.Println(p.curToken)
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if (stmt) != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {

	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatment()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatment() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}

	//returns nil if peektoken is not a identity token after a let , ie. x, if so move to next token
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	//returns nil if peektoken is not assign token after identity token, if so move to next token
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// let x = 5222;
	// 5 -> 2 -> 2 -> 2;
	// TOOO (ignores expression end loop when hit semicolon, denotes end of let statment)
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) peekTokenIs(token token.TokenType) bool {
	return p.peekToken.Type == token
}
func (p *Parser) curTokenIs(token token.TokenType) bool {
	return p.curToken.Type == token
}

func (p *Parser) expectPeek(token token.TokenType) bool {
	if p.peekTokenIs(token) {
		p.nextToken()
		return true
	}
	return false
}
