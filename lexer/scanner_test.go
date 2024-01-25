package lexer

import (
	"DGFC/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "+(;"

	tests := []struct {
		expectedType token.TokenType
		expectedVal  string
	}{
		//{token.ASSIGN, "="},
		{token.ADD, "+"},
		{token.LPAREN, "("},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q",
				i, tt.expectedType, tk.Type)
		}

		if tk.Value != tt.expectedVal {
			t.Fatalf("tests[%d] - value wrong, expected=%q, got=%q",
				i, tt.expectedVal, tk.Value)
		}
	}
}
