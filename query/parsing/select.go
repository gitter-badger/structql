package parsing

import (
	"strings"

	"github.com/s2gatev/structql/query/lexing"
)

// SelectState parses SELECT SQL clauses along with the desired fields.
type SelectState struct{}

func (s *SelectState) Next() []State {
	return []State{
		&FromState{},
	}
}

func (s *SelectState) Parse(result Node, tokenizer *Tokenizer) (Node, bool) {
	if token, _ := tokenizer.ReadToken(); token != lexing.SELECT {
		tokenizer.UnreadToken()
		return result, false
	}

	target := &SelectStatement{}

	// Parse fields.
	for {
		if token, value := tokenizer.ReadToken(); s.isFieldToken(token) {
			target.AddField(s.parseField(value))
		} else {
			return nil, false
		}

		if token, _ := tokenizer.ReadToken(); token != lexing.COMMA {
			tokenizer.UnreadToken()
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

func (s *SelectState) parseField(literal string) *Field {
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
