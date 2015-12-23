package parsing

import (
	"github.com/s2gatev/structql/query/lexing"
)

// OffsetState parses OFFSET SQL clauses along with the value.
type OffsetState struct {
	NextStates []State
}

func (s *OffsetState) Next() []State {
	return s.NextStates
}

func (s *OffsetState) Parse(result Node, tokenizer *Tokenizer) (Node, bool) {
	if target, ok := result.(HasOffset); ok {
		if token, _ := tokenizer.ReadToken(); token != lexing.OFFSET {
			tokenizer.UnreadToken()
			return result, false
		}

		if token, value := tokenizer.ReadToken(); token == lexing.LITERAL {
			target.SetOffset(value)
		} else {
			return result, false
		}

		return result, true
	} else {
		return result, false
	}
}
