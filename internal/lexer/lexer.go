package lexer

import (
	"fmt"
	"rug/internal/token"
	"strings"
	"unicode/utf8"
)

type Lexer struct {
	source  string           // Source code to tokenize
	start   int              // Start position of the current token in source
	pos     int              // Current position in source
	line    int              // Current line number
	linePos int              // Current position in the line
	history runeStack        // History of runes for error recovery
	state   StateFunc        // Current state function
	tokens  chan token.Token // Channel to emit tokens
}

func Tokenize(src string) <-chan token.Token {
	l := &Lexer{
		state:   lexText,
		source:  src,
		tokens:  make(chan token.Token, 255), // Buffered channel to avoid blocking
		start:   0,
		pos:     0,
		line:    1,
		linePos: 1,
		history: newRuneStack(),
	}
	go l.tokenize()
	return l.tokens
}

func (l *Lexer) Current() string {
	return l.source[l.start:l.pos]
}

func (l *Lexer) Emit(t token.TokenType) {
	tok := token.Token{
		Type:    t,
		Literal: l.Current(),
		Line:    l.line,
		Column:  l.linePos,
	}
	l.tokens <- tok
	l.start = l.pos
	l.history.clear()
}

func (l *Lexer) Ignore() {
	l.history.clear()
	l.start = l.pos
}

func (l *Lexer) Peek() rune {
	r := l.Next()
	l.Rewind()
	return r
}

func (l *Lexer) PeekN(n int) rune {
	var r rune
	for i := 0; i < n; i++ {
		r = l.Next()
	}
	for i := 0; i < n; i++ {
		l.Rewind()
	}

	return r
}

const EOF = -1

func (l *Lexer) Rewind() {
	r := l.history.pop()
	if r > EOF {
		size := utf8.RuneLen(r)
		l.pos -= size
		if l.pos < l.start {
			l.pos = l.start
		}
	}
}

func (l *Lexer) Next() rune {
	var (
		r rune
		s int
	)
	str := l.source[l.pos:]
	if len(str) == 0 {
		r, s = EOF, 0
	} else {
		r, s = utf8.DecodeRuneInString(str)
	}
	l.pos += s
	l.history.push(r)

	return r
}

func (l *Lexer) Accept(valid string) bool {
	if strings.ContainsRune(valid, l.Next()) {
		return true
	}
	l.Rewind()
	return false
}

func (l *Lexer) AcceptRun(valid string) {

	for strings.ContainsRune(valid, l.Next()) {
	}
	l.Rewind()
}

func (l *Lexer) Errorf(format string, args ...interface{}) {
	l.tokens <- token.Token{
		Type:    token.ILLEGAL,
		Literal: fmt.Sprintf(format, args...),
		Line:    l.line,
		Column:  l.linePos,
	}
}

func (l *Lexer) tokenize() {
	for state := l.state; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}
