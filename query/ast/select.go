package ast

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

func (ss *Select) SetTarget(name, alias string) {
	ss.TableName = name
	ss.TableAlias = alias
}

func (ss *Select) SetLimit(limit string) {
	ss.Limit = limit
}

func (ss *Select) SetOffset(offset string) {
	ss.Offset = offset
}

func (ss *Select) BuildQuery() string {
	query := ""

	// Build SELECT part.
	fieldsPart := ""
	for index, field := range ss.Fields {
		if index != 0 {
			fieldsPart += ", "
		}
		fieldsPart += field.BuildQuery()
	}
	query += "SELECT " + fieldsPart

	// Build FROM part.
	targetPart := ss.TableName
	if ss.TableAlias != "" {
		targetPart += " " + ss.TableAlias
	}
	query += " FROM " + targetPart

	// Build WHERE part.
	if len(ss.Conditions) > 0 {
		conditionsPart := ""
		for index, condition := range ss.Conditions {
			if index != 0 {
				conditionsPart += " AND "
			}
			conditionsPart += condition.BuildQuery()
		}
		query += " WHERE " + conditionsPart
	}

	// Build LIMIT part.
	if ss.Limit != "" {
		query += " LIMIT " + ss.Limit
	}

	// Build OFFSET part.
	if ss.Offset != "" {
		query += " OFFSET " + ss.Offset
	}
	return query
}
