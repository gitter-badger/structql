package query

import (
	"github.com/s2gatev/structql/query/parsing"
)

type Visitor struct{}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v *Visitor) Visit(node parsing.Node, handle func(parsing.Node) bool) {
	if shouldContinue := handle(node); !shouldContinue {
		return
	}

	switch concrete := node.(type) {
	case *parsing.SelectStatement:
		for _, field := range concrete.Fields {
			v.Visit(field, handle)
		}
		for _, filter := range concrete.Filters {
			v.Visit(filter, handle)
		}
	case *parsing.Field:
	case *parsing.EqualsFilter:
		v.Visit(concrete.Field, handle)
	}
}
