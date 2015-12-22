package parsing

import (
	"io"
	"strings"

	"github.com/s2gatev/structql/query/lexing"
)

type Parser struct {
	lexer  *lexing.Lexer
	buffer struct {
		token  lexing.Token
		value  string
		isFull bool
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{lexer: lexing.NewLexer(r)}
}

func (p *Parser) scan() (lexing.Token, string) {
	if p.buffer.isFull {
		p.buffer.isFull = false
		return p.buffer.token, p.buffer.value
	}

	p.buffer.token, p.buffer.value = p.lexer.NextToken()

	return p.buffer.token, p.buffer.value
}

func (p *Parser) parseField(literal string) *Field {
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

func (p *Parser) parseFilter(literal string) *EqualsFilter {
	literalParts := strings.Split(literal, "=")
	filter := &EqualsFilter{}
	filter.Field = p.parseField(literalParts[0])
	filter.Value = literalParts[1]
	return filter
}

func (p *Parser) Parse() (Node, error) {
	var parser State = &RootState{}
	var result Node
	for {
		if len(parser.Next()) == 0 {
			break
		}
		var parsed bool
		for _, next := range parser.Next() {
			var ok bool
			if result, ok = next.Parse(result, p); ok {
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

func (p *Parser) scanIgnoreWhitespace() (lexing.Token, string) {
	token, value := p.scan()
	if token == lexing.WHITESPACE {
		token, value = p.scan()
	}
	return token, value
}

func (p *Parser) unscan() {
	p.buffer.isFull = true
}
