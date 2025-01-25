package lexer

import (
	"testing"
	"github.com/go-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLOM, ";"},
		{token.EOF, "EOF"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		
		if tt.expectedType != tok.Type {
			t.Fatalf("tests[%d] failed - incorrect tokentype. Expected=%q Actual=%q", i, tt.expectedType, tok.Type)
		}

		if tt.expectedType != tok.Type {
			t.Fatalf("tests[%d] failed - incorrect literal. Expected=%q Actual=%q", i, tt.Literal, tok.Literal)
		}
	}


}