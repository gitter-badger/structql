package parsing

type State interface {
	Parse(Node, *Tokenizer) (Node, bool)
	Next() []State
}
