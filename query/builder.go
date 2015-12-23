package query

import (
	"reflect"
	"strings"

	"github.com/s2gatev/structql/query/ast"
	"github.com/s2gatev/structql/query/parsing"
)

func GenerateQueryFunc(query string, models map[string]reflect.Type) string {
	parser := parsing.NewParser(strings.NewReader(query))
	queryNode, _ := parser.Parse()
	selectQuery := queryNode.(*ast.Select)

	model := models[selectQuery.TableName]

	visitor := NewVisitor()
	visitor.Visit(queryNode, func(node ast.Node) bool {
		if fieldNode, ok := node.(*ast.Field); ok {
			field, _ := model.FieldByName(fieldNode.Name)
			fieldDBColumn := field.Tag.Get("db")
			fieldNode.Name = fieldDBColumn
		} else if selectNode, ok := node.(*ast.Select); ok {
			tableNameField, _ := model.FieldByName("DB_TableName")
			tableName := tableNameField.Tag.Get("db")
			selectNode.TableName = tableName
		}
		return true
	})

	return queryNode.BuildQuery()
}
