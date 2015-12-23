package ast

// HasTarget is an AST node with target.
type HasTarget interface {
	// SetTarget sets the target of the node.
	SetTarget(string, string)
}
