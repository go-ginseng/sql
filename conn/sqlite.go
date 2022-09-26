package conn

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Sqlite connects to SQLite database
func Sqlite(dsn string, silent bool) (*gorm.DB, error) {
	return _createSqliteDB(dsn, silent)
}

// SqliteMemory connects to in-memory SQLite database
func SqliteMemory(shared bool, silent bool) (*gorm.DB, error) {
	dsn := _buildInMemorySqliteDsn(shared)
	return _createSqliteDB(dsn, silent)
}

func _buildInMemorySqliteDsn(shared bool) string {
	dsn := "file::memory:"
	if shared {
		dsn += "?cache=shared"
	}
	return dsn
}

func _createSqliteDB(dsn string, silent bool) (*gorm.DB, error) {
	config := _getConfig(silent)
	return gorm.Open(sqlite.Open(dsn), config)
}
