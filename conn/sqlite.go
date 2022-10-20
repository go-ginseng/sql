package conn

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Connect to SQLite database
func SQLite(filepath string, silent bool) (*gorm.DB, error) {
	return createSQLiteConnection(filepath, silent)
}

// Connect to SQLite database in memory
func SQLiteMemory(shared bool, silent bool) (*gorm.DB, error) {
	dsn := buildInMemorySQLiteDSN(shared)
	return createSQLiteConnection(dsn, silent)
}

func buildInMemorySQLiteDSN(shared bool) string {
	dsn := "file::memory:"
	if shared {
		dsn += "?cache=shared"
	}
	return dsn
}

func createSQLiteConnection(dsn string, silent bool) (*gorm.DB, error) {
	config := getConfig(silent)
	return gorm.Open(sqlite.Open(dsn), config)
}
