package parsing

import (
	"github.com/s2gatev/structql/query/ast"
	"github.com/s2gatev/structql/query/lexing"
)

func SelectTarget(key lexing.Token) func(ast.Node, *Tokenizer) (ast.Node, bool) {
	return func(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
		if target, ok := result.(ast.HasTarget); ok {
			if token, _ := tokenizer.ReadToken(); token != key {
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
}