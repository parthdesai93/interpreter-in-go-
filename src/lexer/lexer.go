package lexer

import (
	"testing"

	"token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedToken  token.TokenType
		expectedLitera string
	}{
		{token.ASSIGN, "="},
	}

	l := New(input)

	for l, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedToken {
			t.Fatal(`tokentype wrong`)
		}

		if tok.Literal !=
	}
}
