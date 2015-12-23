package parsing

import (
	"strings"

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
			if token, value := tokenizer.ReadToken(); token == lexing.LITERAL {
				target.AddFilter(s.parseFilter(value))
			} else {
				panic("WHERE clause must come with conditions.")
			}

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

func (s *WhereState) parseField(literal string) *Field {
	field := &Field{}
	literalParts := strings.Split(literal, ".")
	if len(literalParts) > 1 {
		field.Target = literalParts[0]
		field.Name = literalParts[1]
	} else {
		field.Name = literalParts[0]
	}
	return field
}

func (s *WhereState) parseFilter(literal string) *EqualsFilter {
	literalParts := strings.Split(literal, "=")
	filter := &EqualsFilter{}
	filter.Field = s.parseField(literalParts[0])
	filter.Value = literalParts[1]
	return filter
}
