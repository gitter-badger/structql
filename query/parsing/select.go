package parsing

import (
	"github.com/s2gatev/structql/query/lexing"
)

// SelectState parses SELECT SQL clauses along with the desired fields.
type SelectState struct{}

func (s *SelectState) Next() []State {
	return []State{
		&FromState{},
	}
}

func (s *SelectState) Parse(result Node, p *Parser) (Node, bool) {
	token, _ := p.scanIgnoreWhitespace()
	if token != lexing.SELECT {
		p.unscan()
		return result, false
	}

	target := &SelectStatement{}

	// Parse fields.
	for {
		if token, value := p.scanIgnoreWhitespace(); s.isFieldToken(token) {
			target.AddField(p.parseField(value))
		} else {
			return nil, false
		}

		if token, _ := p.scanIgnoreWhitespace(); token != lexing.COMMA {
			p.unscan()
			break
		}
	}

	return target, true
}

func (s *SelectState) isFieldToken(token lexing.Token) bool {
	switch token {
	case lexing.LITERAL, lexing.ASTERISK:
		return true
	default:
		return false
	}
}
