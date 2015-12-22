package lexing

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// eof is a rune representing the end of file symbol.
var eof = rune(0)

// Lexer is a string tokenizer following SQL syntax rules.
type Lexer struct {
	reader *bufio.Reader
}

// NewLexer creates a new instance of Lexer that scans the tokens of the provided reader.
func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{reader: bufio.NewReader(reader)}
}

// NextToken returns the next SQL token found in the reader.
func (l *Lexer) NextToken() (Token, string) {
	ch := l.read()

	if isWhitespace(ch) {
		l.unread()
		return l.nextWhitespace()
	} else if isLetter(ch) || isDigit(ch) {
		l.unread()
		return l.nextLiteral()
	}

	value := string(ch)
	token := ILLEGAL
	switch ch {
	case eof:
		value = ""
		token = EOF
	case '*':
		token = ASTERISK
	case ',':
		token = COMMA
	}

	return token, value
}

// read reads the next rune from the reader.
func (l *Lexer) read() rune {
	if ch, _, err := l.reader.ReadRune(); err == nil {
		return ch
	}
	return eof
}

// unread puts back the last read rune in the reader.
func (l *Lexer) unread() {
	l.reader.UnreadRune()
}

// nextWhitespace returns the next whitespace token in the reader.
func (l *Lexer) nextWhitespace() (Token, string) {
	var value bytes.Buffer
	value.WriteRune(l.read())

	for {
		if ch := l.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			l.unread()
			break
		} else {
			value.WriteRune(ch)
		}
	}

	return WHITESPACE, value.String()
}

// nextLiteral returns the next literal token in the reader.
func (l *Lexer) nextLiteral() (Token, string) {
	var valueBuffer bytes.Buffer
	valueBuffer.WriteRune(l.read())

	for {
		if ch := l.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && !isSymbol(ch) {
			l.unread()
			break
		} else {
			valueBuffer.WriteRune(ch)
		}
	}

	value := valueBuffer.String()
	token := LITERAL
	switch strings.ToUpper(value) {
	case "SELECT":
		token = SELECT
	case "FROM":
		token = FROM
	case "WHERE":
		token = WHERE
	case "AND":
		token = AND
	case "OR":
		token = OR
	case "LIMIT":
		token = LIMIT
	case "OFFSET":
		token = OFFSET
	}

	return token, value
}
