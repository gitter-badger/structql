package parsing

import (
	"github.com/s2gatev/structql/query/ast"
	"github.com/s2gatev/structql/query/lexing"
)

// UpdateState parses UPDATE SQL clauses along with the desired fields.
type UpdateState struct {
	NextStates []State
}

func (s *UpdateState) Next() []State {
	return s.NextStates
}

func (s *UpdateState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	return SelectTarget(lexing.UPDATE)(&ast.Update{}, tokenizer)
}

func (s *UpdateState) isFieldToken(token lexing.Token) bool {
	switch token {
	case lexing.LITERAL, lexing.ASTERISK:
		return true
	default:
		return false
	}
}
