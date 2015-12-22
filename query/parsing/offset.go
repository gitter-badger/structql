package parsing

import (
	"github.com/s2gatev/structql/query/lexing"
)

// OffsetState parses OFFSET SQL clauses along with the value.
type OffsetState struct{}

func (s *OffsetState) Next() []State {
	return []State{}
}

func (s *OffsetState) Parse(result Node, p *Parser) (Node, bool) {
	if target, ok := result.(HasOffset); ok {
		if token, _ := p.scanIgnoreWhitespace(); token != lexing.OFFSET {
			p.unscan()
			return result, false
		}

		if token, value := p.scanIgnoreWhitespace(); token == lexing.LITERAL {
			target.SetOffset(value)
		} else {
			return result, false
		}

		return result, true
	} else {
		return result, false
	}
}
