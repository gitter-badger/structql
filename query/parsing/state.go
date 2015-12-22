package parsing

type State interface {
	Parse(Node, *Parser) (Node, bool)
	Next() []State
}
