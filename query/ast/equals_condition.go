package ast

type EqualsCondition struct {
	Field *Field
	Value string
}

func (c *EqualsCondition) Assemble() string {
	return c.Field.Assemble() + "=" + c.Value
}
