# structql
[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Conceptually **structql** is a database abstraction layer that goes somewhere between the `database/sql` library and an ORM, offering:
* Fully-generated queries on runtime. Query funcs are generated on build so there is no runtime performance impact.
* Work with struct models instead of database tables and columns.
* Statically-checked queries. Changes to DB schema reflect on the generated queries and cause build issues.

## Models

Models are the interface to the database that a user is supposed to used. They are defined as structs
with fields tagged in a specific way.

```go
//db:model
type User struct {
  DB_TableName struct{} `db:"users"`

  ID int64 `db:"id"`
  Name string `db:"name"`
  Age int `db:"age"`
}
```

## Query functions

Query functions present abstraction over databse queries that works on top of models. Each section shows and example of a query definition and the generated query function from the definition.

### select-where
Query definition:
```go
//db:query GetUserNameByID "SELECT u.Name, u.Age FROM User u WHERE u.ID=?"
```
Generated function:
```go
// GetUserNameByID is a generated query func.
func GetUserNameByID(uID int64) ([]User, error) {
  rows, err := db.Query("SELECT u.name, u.age FROM users u WHERE u.id=?", uID)
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  var results []User
  for rows.Next() {
    resultItem := &User{}
    if err := rows.Scan(&resultItem.Name, &resultItem.Age); err != nil {
      return nil, err
    }
    results = append(results, resultItem)
  }
  return results, nil
}
```

### select-where-limit-offset
Query definition:
```go
//db:query GetUserNameByIDAndAge "SELECT u.Name FROM User u WHERE u.ID=? AND u.Age=? LIMIT 10 OFFSET 20"
```
Generated function:
```go
// GetUserNameByIDAndAge is a generated query func.
func GetUserNameByIDAndAge(uID int64, uAge int) ([]User, error) {
  rows, err := db.Query("SELECT u.name FROM users u WHERE u.id=? AND u.age=? LIMIT 10 OFFSET 20", uID, uAge)
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  var results []User
  for rows.Next() {
    resultItem := &User{}
    if err := rows.Scan(&resultItem.Name); err != nil {
      return nil, err
    }
    results = append(results, resultItem)
  }
  return results, nil
}
```
