package parsing

import (
	"github.com/s2gatev/structql/query/ast"
	"github.com/s2gatev/structql/query/lexing"
)

// SetState parses SET SQL clauses along with their fields.
type SetState struct {
	NextStates []State
}

func (s *SetState) Next() []State {
	return s.NextStates
}

func (s *SetState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	if target, ok := result.(ast.HasFields); ok {
		if token, _ := tokenizer.ReadToken(); token != lexing.SET {
			tokenizer.UnreadToken()
			return result, false
		}

		// Parse WHERE conditions.
		for {
			token, fieldName := tokenizer.ReadToken()
			if token != lexing.LITERAL {
				panic("SET clause must come with conditions.")
			}

			if token, _ := tokenizer.ReadToken(); token != lexing.EQUALS {
				panic("Wrong condition in SET clause.")
			}

			condition := &ast.EqualsCondition{}
			condition.Field = parseField(fieldName)

			token, value := tokenizer.ReadToken()
			if token != lexing.LITERAL && token != lexing.PLACEHOLDER {
				panic("Wrong condition in SET clause.")
			}

			field := parseField(fieldName)
			field.Value = value

			target.AddField(field)

			if token, _ := tokenizer.ReadToken(); token != lexing.COMMA {
				tokenizer.UnreadToken()
				break
			}
		}

		return result, true
	} else {
		return result, false
	}
}
