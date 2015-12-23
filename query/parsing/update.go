package parsing

import (
	"github.com/s2gatev/structql/query/ast"
	"github.com/s2gatev/structql/query/lexing"
)

// UpdateState parses UPDATE SQL clauses along with the target table.
// UPDATE User u ...
type UpdateState struct {
	BaseState
}

func (s *UpdateState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	return SelectTarget(lexing.UPDATE, &ast.Update{}, tokenizer)
}

func (s *UpdateState) isFieldToken(token lexing.Token) bool {
	switch token {
	case lexing.LITERAL, lexing.ASTERISK:
		return true
	default:
		return false
	}
}
