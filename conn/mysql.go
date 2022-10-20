package conn

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect to MySQL database
func MySQL(
	host string,
	port string,
	username string,
	password string,
	database string,
	silent bool,
) (*gorm.DB, error) {
	dsn := buildMySQL_DSN(host, port, username, password, database)
	return createMySQLConnection(dsn, silent)
}

func buildMySQL_DSN(host string, port string, username string, password string, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
}

func createMySQLConnection(dsn string, silent bool) (*gorm.DB, error) {
	config := getConfig(silent)
	return gorm.Open(mysql.Open(dsn), config)
}

func getConfig(silent bool) *gorm.Config {
	config := &gorm.Config{}
	if silent {
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	return config
}
