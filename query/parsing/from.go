package parsing

import (
	"github.com/s2gatev/structql/query/ast"
	"github.com/s2gatev/structql/query/lexing"
)

// FromState parses FROM SQL clauses along with the table name and alias.
type FromState struct {
	BaseState
}

func (s *FromState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	return SelectTarget(lexing.FROM)(result, tokenizer)
}
