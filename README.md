# Go Utility: `sql`

## Installation

```bash
go get -u github.com/nelsonlai-go/sql
```

## Connection

### Connect to MySQL

```go
// import "github.com/nelsonlai-go/sql/conn"

// param: (host, port, username, password, database string, silent bool)
// return: (*sql.DB, error)
db, err := conn.MySql("127.0.0.1", "3306", "root", "password", "database", true)
if err != nil {
    panic(err)
}
```

### Connect to disk-based SQLite

```go
// import "github.com/nelsonlai-go/sql/conn"

// param: (dsn string, silent bool)
// return: (*sql.DB, error)
db, err := conn.Sqlite("./database.db", true)
if err != nil {
    panic(err)
}
```

### Connect to memory-based SQLite

```go
// import "github.com/nelsonlai-go/sql/conn"

// param: (shared, silent bool)
// return: (*sql.DB, error)
db, err := conn.SqliteMemory(":memory:", true)
if err != nil {
    panic(err)
}
```
