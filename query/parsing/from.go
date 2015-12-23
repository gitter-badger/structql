package parsing

import (
	"github.com/s2gatev/structql/query/lexing"
)

// FromState parses FROM SQL clauses along with the table name and alias.
type FromState struct{}

func (s *FromState) Next() []State {
	return []State{
		&WhereState{},
		&LimitState{},
	}
}

func (s *FromState) Parse(result Node, tokenizer *Tokenizer) (Node, bool) {
	if target, ok := result.(HasTarget); ok {
		if token, _ := tokenizer.ReadToken(); token != lexing.FROM {
			tokenizer.UnreadToken()
			return result, false
		}

		// Parse table name.
		token, name := tokenizer.ReadToken()
		if token != lexing.LITERAL {
			panic("FROM clause must come with table name.")
		}

		// Parse table alias.
		alias := ""
		if token, value := tokenizer.ReadToken(); token == lexing.LITERAL {
			alias = value
		} else {
			tokenizer.UnreadToken()
		}

		target.AddTarget(name, alias)

		return result, true
	} else {
		return result, false
	}
}
