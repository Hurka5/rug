package lexer

import (
	"rug/internal/token"
	"strings"
)

type StateFunc func(*Lexer) StateFunc

const (
	NUMS          = ".+-0123456789"
	SPECIAL_CHARS = "(){}[]"
	WHITESPACE    = " \t\r"
)

func lexText(l *Lexer) StateFunc {

	l.AcceptRun(WHITESPACE)
	l.Ignore()

	r := l.Peek()

	// check for EOF
	if r == EOF {
		return nil
	}

	// check for new line
	if r == '\n' {
		l.Next()
		l.Emit(token.NEWLINE)
		return lexText
	}

	// check for comment
	if r == '/' && strings.ContainsRune("/*", l.PeekN(2)) {
		return lexComment
	}

	// check for number
	if strings.ContainsRune(NUMS, r) || r == '.' && strings.ContainsRune(NUMS, l.PeekN(2)) {
		return lexNumber
	}

	return lexIdentifier
}

func lexNumber(l *Lexer) StateFunc {

	tok := token.INT
	digits := "0123456789"

	// Handle sign
	l.Accept("+-")

	// Check for base prefixes
	if l.Accept("0") {
		switch {
		case l.Accept("xX"):
			digits = "0123456789abcdefABCDEF"
		case l.Accept("bB"):
			digits = "01"
		case l.Accept("oO"):
			digits = "01234567"
		default:
			// leading 0 without x/b/o means octal by default (Go-style)
			digits = "01234567"
			l.Rewind()
		}
	}

	l.AcceptRun(digits)

	// Fractional part
	if l.Accept(".") {
		tok = token.FLOAT
		l.AcceptRun(digits)
	}

	// Exponent part
	expType := l.Peek()
	if l.Accept("eEpP") {
		if expType == 'e' || expType == 'E' {
			tok = token.FLOAT
		} else if expType == 'p' || expType == 'P' {
			// only allowed in hex float
			if !strings.ContainsAny(l.Current(), "xX") {
				l.Errorf("'p' exponent requires hexadecimal mantissa")
				l.Ignore()
				return lexText
			}
			tok = token.FLOAT
		}
		l.Accept("+-")
		l.AcceptRun("0123456789")

	}

	// Imaginary suffix
	if l.Accept("i") {
		tok = token.IMAGINARY
	}

	// Emit the token
	l.Emit(tok)

	return lexText
}

func lexComment(l *Lexer) StateFunc {

	l.Accept("/")

	// Check for single-line or multi-line comment
	if l.Accept("/") {
		// Handle single-line comments
		for {
			r := l.Next()
			if r == 10 || r == EOF {
				l.Rewind()
				break
			}
		}
	} else if l.Accept("*") {
		// Handle multi-line comments
		for {
			r := l.Next()
			if r == '*' && l.Peek() == '/' {
				l.Next()
				break
			}
			if r == EOF {
				l.Errorf("Unterminated comment")
				break
			}
		}
	}

	l.Ignore()

	return lexText
}

func lexString(l *Lexer) StateFunc {
	// TODO: Implement string lexing ("asd")
	return lexText
}

func lexRawString(l *Lexer) StateFunc {
	// TODO: Implement raw string lexing (`asd`)
	return lexText
}

func lexChar(l *Lexer) StateFunc {
	// TODO: Implement char lexing ('a')
	return lexText
}

func lexEscape(l *Lexer) StateFunc {
	// TODO: Implement escape sequence lexing (\n, \t, \xFF, \uFFFF)
	return lexText
}

func lexIdentifier(l *Lexer) StateFunc {

	for {
		r := l.Next()
		if strings.ContainsRune(WHITESPACE, r) || r == EOF {
			l.Rewind()
			break
		}
	}

	// Check if the identifier is a keyword
	tok := token.LookupType(l.Current())
	l.Emit(tok)

	return lexText
}
