package query

import (
	"reflect"
	"regexp"
	"strings"
)

func GenerateQueryFunc(query string, models map[string]reflect.Type) string {
	targetRegexp := regexp.MustCompile(`FROM\s+(?P<model>[[:alnum:]]+)\s+(?P<alias>[[:alnum:]]+)`)
	match := targetRegexp.FindStringSubmatch(query)
	groups := make(map[string]string)
	for i, name := range targetRegexp.SubexpNames() {
		if i != 0 {
			groups[name] = match[i]
		}
	}

	target := models[groups["model"]]

	fieldRegexp := regexp.MustCompile(groups["alias"] + `\.(?P<field>[[:alnum:]]+)`)
	result := fieldRegexp.ReplaceAllFunc([]byte(query), func(match []byte) []byte {
		matchParts := strings.Split(string(match), ".")
		modelAlias, fieldName := matchParts[0], matchParts[1]
		field, _ := target.FieldByName(fieldName)
		fieldColumnName := field.Tag.Get("db")
		return []byte(modelAlias + "." + fieldColumnName)
	})

	result = targetRegexp.ReplaceAllFunc(result, func(match []byte) []byte {
		tableNameField, _ := target.FieldByName("DB_TableName")
		tableName := tableNameField.Tag.Get("db")
		return []byte("FROM " + tableName + " " + groups["alias"])
	})

	return string(result)
}
