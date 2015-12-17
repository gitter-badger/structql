package query_test

import (
	"reflect"
	"testing"

	. "github.com/s2gatev/structql/query"
)

func testQuery(
	t *testing.T,
	query string,
	expected string,
	modelObjects ...interface{}) {

	models := map[string]reflect.Type{}
	for _, modelObject := range modelObjects {
		modelType := reflect.TypeOf(modelObject)
		models[modelType.Name()] = modelType
	}

	actual := GenerateQueryFunc(query, models)
	if actual != expected {
		t.Errorf("Different queries.\n"+
			"Expected: %v\n"+
			"Actual: %v\n", expected, actual)
	}
}

func TestSimpleSelectQuery(t *testing.T) {
	type User struct {
		DB_TableName struct{} `db:"users"`

		ID   int64  `db:"id"`
		Name string `db:"name"`
	}

	testQuery(t,
		"SELECT u.Name FROM User u WHERE u.ID=?",
		"SELECT u.name FROM users u WHERE u.id=?",
		User{})
}

func TestSelectMultipleFieldsMap(t *testing.T) {
	type User struct {
		DB_TableName struct{} `db:"users"`

		ID   int64  `db:"id"`
		Name string `db:"name"`
		Age  int    `db:"age"`
	}

	testQuery(t,
		"SELECT u.Name, u.Age FROM User u WHERE u.ID=?",
		"SELECT u.name, u.age FROM users u WHERE u.id=?",
		User{})
}

func TestSelectMultipleFieldsFilter(t *testing.T) {
	type User struct {
		DB_TableName struct{} `db:"users"`

		ID   int64  `db:"id"`
		Name string `db:"name"`
		Age  int    `db:"age"`
	}

	testQuery(t,
		"SELECT u.Name FROM User u WHERE u.ID=? AND u.Age=?",
		"SELECT u.name FROM users u WHERE u.id=? AND u.age=?",
		User{})
}

func TestSelectLimit(t *testing.T) {
	type User struct {
		DB_TableName struct{} `db:"users"`

		ID   int64  `db:"id"`
		Name string `db:"name"`
		Age  int    `db:"age"`
	}

	testQuery(t,
		"SELECT u.Name FROM User u WHERE u.ID=? AND u.Age=? LIMIT 10",
		"SELECT u.name FROM users u WHERE u.id=? AND u.age=? LIMIT 10",
		User{})
}

func TestSelectLimitOffset(t *testing.T) {
	type User struct {
		DB_TableName struct{} `db:"users"`

		ID   int64  `db:"id"`
		Name string `db:"name"`
		Age  int    `db:"age"`
	}

	testQuery(t,
		"SELECT u.Name FROM User u WHERE u.ID=? AND u.Age=? LIMIT 10 OFFSET 20",
		"SELECT u.name FROM users u WHERE u.id=? AND u.age=? LIMIT 10 OFFSET 20",
		User{})
}
