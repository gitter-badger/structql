package parsing

import (
	"github.com/s2gatev/structql/query/ast"
	"github.com/s2gatev/structql/query/lexing"
)

// SelectState parses SELECT SQL clauses along with the desired fields.
// SELECT u.Name, u.Age ...
type SelectState struct {
	BaseState
}

func (s *SelectState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	if token, _ := tokenizer.ReadToken(); token != lexing.SELECT {
		tokenizer.UnreadToken()
		return result, false
	}

	target := &ast.Select{}

	// Parse fields.
	for {
		if token, value := tokenizer.ReadToken(); s.isFieldToken(token) {
			target.AddField(parseField(value))
		} else {
			return nil, false
		}

		if token, _ := tokenizer.ReadToken(); token != lexing.COMMA {
			tokenizer.UnreadToken()
			break
		}
	}

	return target, true
}

func (s *SelectState) isFieldToken(token lexing.Token) bool {
	switch token {
	case lexing.LITERAL, lexing.ASTERISK:
		return true
	default:
		return false
	}
}
