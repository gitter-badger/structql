package parsing

import (
	"strings"

	"github.com/s2gatev/structql/query/ast"
)

// parseField parses field extracting its name and target.
func parseField(literal string) *ast.Field {
	field := &ast.Field{}
	literalParts := strings.Split(literal, ".")
	if len(literalParts) > 1 {
		field.Target = literalParts[0]
		field.Name = literalParts[1]
	} else {
		field.Name = literalParts[0]
	}
	return field
}
