package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	letToken := token.Token{Type: token.LET, Literal: "let"}
	varToken := token.Token{Type: token.LET, Literal: "myVar"}
	letIdentifier := &Identifier{
		Token: varToken,
		Value: "myVar",
	}
	anotherIdentifier := &Identifier{
		Token: varToken,
		Value: "anotherVar",
	}
	letStatement := &LetStatement{
		Token: letToken,
		Name:  letIdentifier,
		Value: anotherIdentifier,
	}
	program := &Program{
		Statements: []Statement{
			letStatement,
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
