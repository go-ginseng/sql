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

## Encryption

### AES encrypt and decrypt

```go
// import "github.com/nelsonlai-go/sql/encrypt"

encryption, err := encrypt.NewAesEncrypt("aes key")
data := "data to encrypt"

// param: (plainText string)
// return: (string, error)
encrypted, err := encryption.Encrypt(data)
if err != nil {
    panic(err)
}

// param: (cipherText string)
// return: (string, error)
decrypted, err := encryption.Decrypt(encrypted)
if err != nil {
    panic(err)
}
```

## Clause Builder

```go
// import "github.com/nelsonlai-go/sql"

// id = ?
cls := sql.Eq("id", 1)

// id <> ?
cls := sql.Neq("id", 1)

// id > ?
cls := sql.Gt("id", 1)

// id >= ?
cls := sql.Gte("id", 1)

// id < ?
cls := sql.Lt("id", 1)

// id <= ?
cls := sql.Lte("id", 1)

// id IN (?, ?, ?)
cls := sql.In("id", []int{1, 2, 3})

// id NOT IN (?, ?, ?)
cls := sql.Nin("id", []int{1, 2, 3})

// id LIKE ?
cls := sql.Lk("id", "value")

// id NOT LIKE ?
cls := sql.Nlk("id", "value")

// id IS NULL
cls := sql.Null("id")

// id IS NOT NULL
cls := sql.NotNull("id")

// id BETWEEN ? AND ?
cls := sql.Between("id", 1, 10)

// (id = ? AND name <> ?)
cls := sql.And(sql.Eq("id", 1), sql.Neq("name", "name"))

// (id = ? OR name <> ?)
cls := sql.Or(sql.Eq("id", 1), sql.Neq("name", "name"))

// build clause and use it
stm, values := cls.Build()
db.Where(stm, values...).Find(&data)
```
