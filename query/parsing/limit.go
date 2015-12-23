package parsing

import (
	"github.com/s2gatev/structql/query/ast"
	"github.com/s2gatev/structql/query/lexing"
)

// LimitState parses LIMIT SQL clauses along with the value.
type LimitState struct {
	NextStates []State
}

func (ls *LimitState) Next() []State {
	return ls.NextStates
}

func (ls *LimitState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	if target, ok := result.(ast.HasLimit); ok {
		if token, _ := tokenizer.ReadToken(); token != lexing.LIMIT {
			tokenizer.UnreadToken()
			return result, false
		}

		if token, value := tokenizer.ReadToken(); token == lexing.LITERAL {
			target.SetLimit(value)
		} else {
			return nil, false
		}

		return result, true
	} else {
		return nil, false
	}
}
