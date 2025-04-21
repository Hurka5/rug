package token

import (
	"strconv"
)

type TokenType int

func (tt TokenType) String() string {
	return typeLookupTable[tt]
}

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

func (t Token) String() string {
	return t.Type.String() + "(" + t.Literal + ")" + "@" + strconv.Itoa(t.Line) + ":" + strconv.Itoa(t.Column)
}

const (
	ILLEGAL TokenType = iota
	EOF
	NEWLINE
	COMMENT
	IDENTIFIER
	INT
	FLOAT
	IMAGINARY
	STRING
	PLUS
	MINUS
	ASTERISK
	SLASH
	LPAREN
	RPAREN
	COMMA
	SEMICOLON
	ASSIGN
	EQUAL
	NOT_EQUAL
	GREATER
	LESS
	GREATER_EQUAL
	LESS_EQUAL
	IF
	ELSE
)

func LookupType(s string) TokenType {
	if tok, ok := lookupTable[s]; ok {
		return tok
	}
	return IDENTIFIER
}

var lookupTable = map[string]TokenType{
	"if":   IF,
	"else": ELSE,
	"+":    PLUS,
	"-":    MINUS,
	"*":    ASTERISK,
	"/":    SLASH,
	"(":    LPAREN,
	")":    RPAREN,
	",":    COMMA,
	";":    SEMICOLON,
	"=":    ASSIGN,
	"==":   EQUAL,
	"!=":   NOT_EQUAL,
	">":    GREATER,
	"<":    LESS,
	">=":   GREATER_EQUAL,
	"<=":   LESS_EQUAL,
}

var typeLookupTable = map[TokenType]string{
	ILLEGAL:       "ILLEGAL",
	EOF:           "EOF",
	NEWLINE:       "NEWLINE",
	COMMENT:       "COMMENT",
	IDENTIFIER:    "IDENTIFIER",
	INT:           "INT",
	FLOAT:         "FLOAT",
	STRING:        "STRING",
	PLUS:          "PLUS",
	MINUS:         "MINUS",
	ASTERISK:      "ASTERISK",
	SLASH:         "SLASH",
	LPAREN:        "LPAREN",
	RPAREN:        "RPAREN",
	COMMA:         "COMMA",
	SEMICOLON:     "SEMICOLON",
	ASSIGN:        "ASSIGN",
	EQUAL:         "EQUAL",
	NOT_EQUAL:     "NOT_EQUAL",
	GREATER:       "GREATER",
	LESS:          "LESS",
	GREATER_EQUAL: "GREATER_EQUAL",
	LESS_EQUAL:    "LESS_EQUAL",
	IF:            "IF",
	ELSE:          "ELSE",
}
