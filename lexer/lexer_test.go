package lexer

import (
	"testing"

	"github.com/arjunsingh14/go-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;

let ten = 10;

let add = fn(x, y) {
	x + y;
}

let result = add(five, ten)
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LBRACE, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tt.expectedType != tok.Type {
			t.Fatalf("tests[%d] failed - incorrect tokentype. Expected=%q Actual=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] failed - incorrect literal. Expected=%q Actual=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
