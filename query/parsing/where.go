package parsing

import (
	"github.com/s2gatev/structql/query/lexing"
)

// WhereState parses WHERE SQL clauses along with their conditions.
type WhereState struct{}

func (s *WhereState) Next() []State {
	return []State{
		&LimitState{},
	}
}

func (s *WhereState) Parse(result Node, tokenizer *Tokenizer) (Node, bool) {
	if target, ok := result.(HasFilters); ok {
		if token, _ := tokenizer.ReadToken(); token != lexing.WHERE {
			tokenizer.UnreadToken()
			return result, false
		}

		// Parse WHERE conditions.
		for {
			token, fieldName := tokenizer.ReadToken()
			if token != lexing.LITERAL {
				panic("WHERE clause must come with conditions.")
			}

			if token, _ := tokenizer.ReadToken(); token != lexing.EQUALS {
				panic("Wrong condition in WHERE clause.")
			}

			filter := &EqualsFilter{}
			filter.Field = parseField(fieldName)

			token, value := tokenizer.ReadToken()
			if token != lexing.LITERAL && token != lexing.PLACEHOLDER {
				panic("Wrong condition in WHERE clause.")
			}

			filter.Value = value
			target.AddFilter(filter)

			if token, _ := tokenizer.ReadToken(); token != lexing.AND {
				tokenizer.UnreadToken()
				break
			}
		}

		return result, true
	} else {
		return result, false
	}
}
