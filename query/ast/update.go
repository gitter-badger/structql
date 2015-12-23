package ast

// Update represents an Update SQL query AST node.
type Update struct {
	Fields     []*Field
	TableName  string
	TableAlias string
	Conditions []*EqualsCondition
}

func (u *Update) AddField(field *Field) {
	u.Fields = append(u.Fields, field)
}

func (u *Update) AddCondition(condition *EqualsCondition) {
	u.Conditions = append(u.Conditions, condition)
}

func (u *Update) SetTarget(name, alias string) {
	u.TableName = name
	u.TableAlias = alias
}

func (u *Update) BuildQuery() string {
	return ""
}
