package database

import (
	"fmt"

	"github.com/dedihartono801/go-clean-architecture-v2/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(config *config.Config) *gorm.DB {
	dsn := getDataSourceName(config)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return database
}

func getDataSourceName(config *config.Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseName,
	)
}
