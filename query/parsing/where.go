package parsing

import (
	"github.com/s2gatev/structql/query/ast"
	"github.com/s2gatev/structql/query/lexing"
)

// WhereState parses WHERE SQL clauses along with their conditions.
// ... WHERE u.Age=? ...
type WhereState struct {
	BaseState
}

func (s *WhereState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	if target, ok := result.(ast.HasConditions); ok {
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

			condition := &ast.EqualsCondition{}
			condition.Field = parseField(fieldName)

			token, value := tokenizer.ReadToken()
			if token != lexing.LITERAL && token != lexing.PLACEHOLDER {
				panic("Wrong condition in WHERE clause.")
			}

			condition.Value = value
			target.AddCondition(condition)

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
