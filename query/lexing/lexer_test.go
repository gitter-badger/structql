package lexing_test

import (
	"strings"
	"testing"

	. "github.com/s2gatev/structql/query/lexing"
)

type lexerTest struct {
	statement string
	token     Token
	literal   string
}

var lexerTests = []lexerTest{
	{``, EOF, ""},
	{`#`, ILLEGAL, `#`},
	{` `, WHITESPACE, " "},
	{"\t", WHITESPACE, "\t"},
	{"\n", WHITESPACE, "\n"},
	{`*`, ASTERISK, "*"},
	{`user`, LITERAL, `user`},
	{`u.name`, LITERAL, `u.name`},
	{`FROM`, FROM, "FROM"},
	{`SELECT`, SELECT, "SELECT"},
	{`WHERE`, WHERE, "WHERE"},
	{`AND`, AND, "AND"},
	{`OR`, OR, "OR"},
	{`LIMIT`, LIMIT, "LIMIT"},
	{`OFFSET`, OFFSET, "OFFSET"},
}

func TestLexer(t *testing.T) {
	for _, test := range lexerTests {
		lexer := NewLexer(strings.NewReader(test.statement))
		token, literal := lexer.NextToken()
		if test.token != token {
			t.Errorf("Token for %v is not correct.\n"+
				"Expected: %v\n"+
				"Actual: %v\n", test.statement, test.token, token)
		}
		if test.literal != literal {
			t.Errorf("Literal for %v is not correct.\n"+
				"Expected: %v\n"+
				"Actual: %v\n", test.statement, test.literal, literal)
		}
	}
}
