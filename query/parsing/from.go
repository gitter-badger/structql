package parsing

import (
	"github.com/s2gatev/structql/query/lexing"
)

// FromState parses FROM SQL clauses along with the table name and alias.
type FromState struct {
	NextStates []State
}

func (s *FromState) Next() []State {
	return s.NextStates
}

func (s *FromState) Parse(result Node, tokenizer *Tokenizer) (Node, bool) {
	return SelectTarget(lexing.FROM)(result, tokenizer)
}
