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
