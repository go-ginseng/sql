package conn

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MySQL connects to MySQL database
func MySql(
	host string,
	port string,
	username string,
	password string,
	database string,
	silent bool,
) (*gorm.DB, error) {
	dsn := _buildMySqlDsn(host, port, username, password, database)
	return _createMySqlDB(dsn, silent)
}

func _buildMySqlDsn(host string, port string, username string, password string, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
}

func _createMySqlDB(dsn string, silent bool) (*gorm.DB, error) {
	config := _getConfig(silent)
	return gorm.Open(mysql.Open(dsn), config)
}

func _getConfig(silent bool) *gorm.Config {
	config := &gorm.Config{}
	if silent {
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	return config
}
