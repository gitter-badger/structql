package parsing_test

import (
	"reflect"
	"strings"
	"testing"

	. "github.com/s2gatev/structql/query/parsing"
)

func testParseCorrectQuery(t *testing.T,
	query string,
	expectedNode Node) {

	actualNode, err := NewParser(strings.NewReader(query)).Parse()
	if err != nil || !reflect.DeepEqual(expectedNode, actualNode) {
		t.Errorf("Parsed node is not correct.\n"+
			"Expected: %+v\n"+
			"Actual: %+v", expectedNode, actualNode)
	}
}

func testParseWrongQuery(t *testing.T,
	query string,
	expectedError string) {

	_, err := NewParser(strings.NewReader(query)).Parse()
	if err == nil || expectedError != err.Error() {
		t.Errorf("Error is not correct.\n"+
			"Expected: %+v\n"+
			"Actual: %+v", expectedError, err.Error())
	}
}

func TestSelectSingleFieldFrom(t *testing.T) {
	testParseCorrectQuery(t,
		`SELECT Name FROM User`,
		&SelectStatement{
			Fields: []*Field{
				&Field{Name: "Name"},
			},
			TableName: "User",
		})
}

func TestSelectMultipleFieldsFrom(t *testing.T) {
	testParseCorrectQuery(t,
		`SELECT Name, Location, Age FROM User`,
		&SelectStatement{
			Fields: []*Field{
				&Field{Name: "Name"}, &Field{Name: "Location"}, &Field{Name: "Age"},
			},
			TableName: "User",
		})
}

func TestSelectAllFieldsFrom(t *testing.T) {
	testParseCorrectQuery(t,
		`SELECT * FROM User`,
		&SelectStatement{
			Fields: []*Field{
				&Field{Name: "*"},
			},
			TableName: "User",
		})
}

func TestSelectFromWithAlias(t *testing.T) {
	testParseCorrectQuery(t,
		`SELECT u.Name, u.Location, u.Age FROM User u`,
		&SelectStatement{
			Fields: []*Field{
				&Field{Target: "u", Name: "Name"},
				&Field{Target: "u", Name: "Location"},
				&Field{Target: "u", Name: "Age"},
			},
			TableName:  "User",
			TableAlias: "u",
		})
}

func TestSelectFromLimitWithAlias(t *testing.T) {
	testParseCorrectQuery(t,
		`SELECT u.Name, u.Location, u.Age FROM User u LIMIT 10`,
		&SelectStatement{
			Fields: []*Field{
				&Field{Target: "u", Name: "Name"},
				&Field{Target: "u", Name: "Location"},
				&Field{Target: "u", Name: "Age"},
			},
			Limit:      "10",
			TableName:  "User",
			TableAlias: "u",
		})
}

func TestSelectFromLimitOffsetWithAlias(t *testing.T) {
	testParseCorrectQuery(t,
		`SELECT u.Name, u.Location, u.Age FROM User u LIMIT 10 OFFSET 20`,
		&SelectStatement{
			Fields: []*Field{
				&Field{Target: "u", Name: "Name"},
				&Field{Target: "u", Name: "Location"},
				&Field{Target: "u", Name: "Age"},
			},
			Limit:      "10",
			Offset:     "20",
			TableName:  "User",
			TableAlias: "u",
		})
}

func TestSelectFromWhereWithAlias(t *testing.T) {
	testParseCorrectQuery(t,
		`SELECT u.Name FROM User u WHERE u.Age=21`,
		&SelectStatement{
			Fields: []*Field{
				&Field{Target: "u", Name: "Name"},
			},
			Filters: []*EqualsFilter{
				&EqualsFilter{
					Field: &Field{Target: "u", Name: "Age"},
					Value: "21",
				},
			},
			TableName:  "User",
			TableAlias: "u",
		})
}
