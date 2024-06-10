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

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	// fmt.Println(p.curToken)
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

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TOOO (ignores expression end loop when hit semicolon, denotes end of return statement)
	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt

}

func (p *Parser) parseStatement() ast.Statement {

	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
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
	// TOOO (ignores expression end loop when hit semicolon, denotes end of let statement)
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
	p.peekError(token)
	return false
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
