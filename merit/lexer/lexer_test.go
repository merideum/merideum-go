package lexer

import (
	"testing"

	"github.com/merideum/merideum-go/merit/token"
)

func TestNextToken(t *testing.T) {
	input := `const test = 123`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CONST, "const"},
		{token.IDENTIFIER, "test"},
		{token.ASSIGN, "="},
		{token.INTEGER, "123"},
		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
