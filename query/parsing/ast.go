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

type HasFilters interface {
	AddFilter(*EqualsFilter)
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

type EqualsFilter struct {
	Field *Field
	Value string
}

func (ef *EqualsFilter) Assemble() string {
	filter := ""
	filter += ef.Field.Assemble()
	filter += "="
	filter += ef.Value
	return filter
}

// SelectStatement represents a SQL SELECT statement.
type SelectStatement struct {
	Fields     []*Field
	TableName  string
	TableAlias string
	Filters    []*EqualsFilter
	Limit      string
	Offset     string
}

func (ss *SelectStatement) AddField(field *Field) {
	ss.Fields = append(ss.Fields, field)
}

func (ss *SelectStatement) AddFilter(filter *EqualsFilter) {
	ss.Filters = append(ss.Filters, filter)
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

	if len(ss.Filters) > 0 {
		filters := ""
		for index, filter := range ss.Filters {
			if index != 0 {
				filters += " AND "
			}
			filters += filter.Assemble()
		}
		query += " WHERE " + filters
	}

	if ss.Limit != "" {
		query += " LIMIT " + ss.Limit
	}
	if ss.Offset != "" {
		query += " OFFSET " + ss.Offset
	}
	return query
}
