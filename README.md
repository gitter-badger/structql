# structql

[![Join the chat at https://gitter.im/s2gatev/structql](https://badges.gitter.im/s2gatev/structql.svg)](https://gitter.im/s2gatev/structql?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

**structql** is a database abstraction layer that sits somewhere between the `database/sql` library and an ORM, offering:
* Fully-generated queries on runtime. Query funcs are generated on build so there is no runtime performance impact.
* Work with struct models instead of database tables and columns.
* Statically-checked queries. Changes to DB schema reflect on the generated queries and cause build issues.

```go
//db:model
type User struct {
	DB_TableName struct{} `db:"users"`

	ID   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}


func PrintUserInfo(userID int64) {
  //db:query GetUserNameAgeByID "SELECT u.Name, u.Age FROM User u WHERE u.ID=?"
  if user, err := queries.GetUserNameAgeByID(userID); err != nil {
    fmt.Println(user.Name)
    fmt.Println(user.Age)
  }
}
```
