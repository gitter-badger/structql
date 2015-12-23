package query

import (
	"github.com/s2gatev/structql/query/ast"
)

type Visitor struct{}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v *Visitor) Visit(node ast.Node, handle func(ast.Node) bool) {
	if shouldContinue := handle(node); !shouldContinue {
		return
	}

	switch concrete := node.(type) {
	case *ast.Select:
		for _, field := range concrete.Fields {
			v.Visit(field, handle)
		}
		for _, condition := range concrete.Conditions {
			v.Visit(condition, handle)
		}
	case *ast.Field:
	case *ast.EqualsCondition:
		v.Visit(concrete.Field, handle)
	}
}
