package parsing

import (
	"github.com/s2gatev/structql/query/lexing"
)

// LimitState parses LIMIT SQL clauses along with the value.
type LimitState struct{}

func (ls *LimitState) Next() []State {
	return []State{
		&OffsetState{},
	}
}

func (ls *LimitState) Parse(result Node, p *Parser) (Node, bool) {
	if target, ok := result.(HasLimit); ok {
		if token, _ := p.scanIgnoreWhitespace(); token != lexing.LIMIT {
			p.unscan()
			return result, false
		}

		if token, value := p.scanIgnoreWhitespace(); token == lexing.LITERAL {
			target.SetLimit(value)
		} else {
			return nil, false
		}

		return result, true
	} else {
		return nil, false
	}
}
