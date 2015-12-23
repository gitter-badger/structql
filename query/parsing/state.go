package parsing

import (
	"github.com/s2gatev/structql/query/ast"
)

type State interface {
	Parse(ast.Node, *Tokenizer) (ast.Node, bool)
	Next() []State
}
