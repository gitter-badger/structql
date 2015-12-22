package query

import (
	"reflect"
	"strings"

	"github.com/s2gatev/structql/query/parsing"
)

func GenerateQueryFunc(query string, models map[string]reflect.Type) string {
	parser := parsing.NewParser(strings.NewReader(query))
	queryNode, _ := parser.Parse()
	selectQuery := queryNode.(*parsing.SelectStatement)

	model := models[selectQuery.TableName]

	visitor := NewVisitor()
	visitor.Visit(queryNode, func(node parsing.Node) bool {
		if fieldNode, ok := node.(*parsing.Field); ok {
			field, _ := model.FieldByName(fieldNode.Name)
			fieldDBColumn := field.Tag.Get("db")
			fieldNode.Name = fieldDBColumn
		} else if selectNode, ok := node.(*parsing.SelectStatement); ok {
			tableNameField, _ := model.FieldByName("DB_TableName")
			tableName := tableNameField.Tag.Get("db")
			selectNode.TableName = tableName
		}
		return true
	})

	return queryNode.Assemble()
}
