package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program does not have enough statements, got=%d", len(program.Statements))
	}

	stmt := program.Statements[0]
	expStmt, ok := stmt.(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statemetnt is not a expression statement, got=%T", stmt)
	}

	literal, ok := expStmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("Expression is not a ast.Identifier type, got=%T", expStmt.Expression)
	}

	if literal.Value != 5 {
		t.Fatalf("Expression value is not 5 got=%d", literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Fatalf("Expression value is not 5 got=%s", literal.TokenLiteral())
	}

}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program got enough statements, got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement is not ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("statement is not an ast.Identifier, got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Fatalf("identity value not %s, got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("identity token literal not %s, got=%s", "foobar", ident.TokenLiteral())
	}

}

func TestReturnStatements(t *testing.T) {
	input := `return 5;
	return a;
	return  83838;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if program == nil {
		t.Fatalf("Parse Program return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not ast.ReturnStatement, got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral is not 'return', got =%s", returnStmt.TokenLiteral())
		}
	}

}

func TestLetStatements(t *testing.T) {
	input := `let x = 5;
	let y = 10;
	let foobar = 83838;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if program == nil {
		t.Fatalf("Parse Program return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	test := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range test {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()

}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let', got=%s", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatment. got=%s", stmt)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Vakue not %s, got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not %s, got=%s", name, letStmt.Name.TokenLiteral())
	}

	return true
}
