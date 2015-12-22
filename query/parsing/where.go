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

func (s *WhereState) Parse(result Node, p *Parser) (Node, bool) {
	if target, ok := result.(HasFilters); ok {
		if token, _ := p.scanIgnoreWhitespace(); token != lexing.WHERE {
			p.unscan()
			return result, false
		}

		// Parse WHERE conditions.
		for {
			if token, value := p.scanIgnoreWhitespace(); token == lexing.LITERAL {
				target.AddFilter(p.parseFilter(value))
			} else {
				panic("WHERE clause must come with conditions.")
			}

			if token, _ := p.scanIgnoreWhitespace(); token != lexing.AND {
				p.unscan()
				break
			}
		}

		return result, true
	} else {
		return result, false
	}
}
