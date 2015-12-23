package parsing

import (
	"fmt"
)

type Node interface {
	Assemble() string
}

type HasFields interface {
	AddField(*Field)
}

type HasConditions interface {
	AddCondition(*EqualsCondition)
}

type HasTarget interface {
	AddTarget(string, string)
}

type HasLimit interface {
	SetLimit(string)
}

type HasOffset interface {
	SetOffset(string)
}

type Field struct {
	Target string
	Name   string
}

func (f *Field) Assemble() string {
	field := ""
	if f.Target != "" {
		field += f.Target + "."
	}
	field += f.Name
	return field
}

type EqualsCondition struct {
	Field *Field
	Value string
}

func (c *EqualsCondition) Assemble() string {
	return c.Field.Assemble() + "=" + c.Value
}

// SelectStatement represents a SQL SELECT statement.
type SelectStatement struct {
	Fields     []*Field
	TableName  string
	TableAlias string
	Conditions []*EqualsCondition
	Limit      string
	Offset     string
}

func (ss *SelectStatement) AddField(field *Field) {
	ss.Fields = append(ss.Fields, field)
}

func (ss *SelectStatement) AddCondition(condition *EqualsCondition) {
	ss.Conditions = append(ss.Conditions, condition)
}

func (ss *SelectStatement) AddTarget(name, alias string) {
	ss.TableName = name
	ss.TableAlias = alias
}

func (ss *SelectStatement) SetLimit(limit string) {
	ss.Limit = limit
}

func (ss *SelectStatement) SetOffset(offset string) {
	ss.Offset = offset
}

func (ss *SelectStatement) Assemble() string {
	fields := ""
	for index, field := range ss.Fields {
		if index != 0 {
			fields += ", "
		}
		fields += field.Assemble()
	}

	target := ss.TableName
	if ss.TableAlias != "" {
		target += " " + ss.TableAlias
	}

	query := fmt.Sprintf("SELECT %v FROM %v", fields, target)

	if len(ss.Conditions) > 0 {
		conditions := ""
		for index, condition := range ss.Conditions {
			if index != 0 {
				conditions += " AND "
			}
			conditions += condition.Assemble()
		}
		query += " WHERE " + conditions
	}

	if ss.Limit != "" {
		query += " LIMIT " + ss.Limit
	}
	if ss.Offset != "" {
		query += " OFFSET " + ss.Offset
	}
	return query
}
