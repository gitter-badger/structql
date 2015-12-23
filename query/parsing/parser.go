package parsing

import (
	"io"

	"github.com/s2gatev/structql/query/ast"
)

// Parser parses SQL query into AST.
type Parser struct {
	tokenizer *Tokenizer
}

// NewParser creates a Parser instance for the provided query.
func NewParser(queryReader io.Reader) *Parser {
	return &Parser{tokenizer: NewTokenizer(queryReader)}
}

// Parse parses the query into a Node.
func (p *Parser) Parse() (ast.Node, error) {
	var parser State = &RootState{}
	var result ast.Node
	for {
		if len(parser.Next()) == 0 {
			break
		}
		var parsed bool
		for _, next := range parser.Next() {
			var ok bool
			if result, ok = next.Parse(result, p.tokenizer); ok {
				parser = next
				parsed = true
				break
			}
		}
		if parsed == false {
			break
		}
	}
	return result, nil
}
