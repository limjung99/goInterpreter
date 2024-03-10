package lexer

import (
	"goLang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ","},
		{token.EOF, ""},
	}

	l := New(input)

	for idx, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			// test fail
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", idx, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedType {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", idx, tt.expectedLiteral, tok.Literal)
		}
	}

}
