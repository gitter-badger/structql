package ast

import (
	"fmt"
)

// Select represents a SQL SELECT statement.
type Select struct {
	Fields     []*Field
	TableName  string
	TableAlias string
	Conditions []*EqualsCondition
	Limit      string
	Offset     string
}

func (ss *Select) AddField(field *Field) {
	ss.Fields = append(ss.Fields, field)
}

func (ss *Select) AddCondition(condition *EqualsCondition) {
	ss.Conditions = append(ss.Conditions, condition)
}

func (ss *Select) AddTarget(name, alias string) {
	ss.TableName = name
	ss.TableAlias = alias
}

func (ss *Select) SetLimit(limit string) {
	ss.Limit = limit
}

func (ss *Select) SetOffset(offset string) {
	ss.Offset = offset
}

func (ss *Select) Assemble() string {
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
