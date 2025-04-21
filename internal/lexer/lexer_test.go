package lexer

import (
	"rug/internal/token"
	"testing"
)

func TestValidTokenLexing(t *testing.T) {
	var tests = map[string][]token.TokenType{
		".12":      {token.FLOAT},
		"0.12":     {token.FLOAT},
		"123":      {token.INT},
		"0b1":      {token.INT},
		"0o12":     {token.INT},
		"001":      {token.INT},
		"0x12":     {token.INT},
		"01":       {token.INT},
		"12i":      {token.IMAGINARY},
		"12.3e2":   {token.FLOAT},
		"12.3e+2":  {token.FLOAT},
		"12.3e-2":  {token.FLOAT},
		"12.3e+2i": {token.IMAGINARY},
		"0x123p-2": {token.FLOAT},
		"0x123p+2": {token.FLOAT},
		"asd":      {token.IDENTIFIER},
		"if":       {token.IF},
		"a b":      {token.IDENTIFIER, token.IDENTIFIER},
		"asd //this is a single line comment \n dsa":  {token.IDENTIFIER, token.NEWLINE, token.IDENTIFIER},
		"asd /*this is a multi line comment \n \n */": {token.IDENTIFIER},
	}

	for input, expected := range tests {
		l := Tokenize(input)

		var found []token.TokenType
		for tok := range l {
			found = append(found, tok.Type)
		}

		if len(found) != len(expected) {
			t.Errorf("INPUT: %s\n Expected %d tokens, got %d", input, len(expected), len(found))
			for _, tok := range found {
				t.Errorf("Found %s", tok.String())
			}
		} else {
			for i, tok := range found {
				if tok != expected[i] {
					t.Errorf("INPUT: %s\n Expected %s, got %s", input, expected[i].String(), tok.String())
				}
			}
		}

	}
}
