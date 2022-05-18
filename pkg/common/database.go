package common

import (
	"fmt"

	"github.com/arashi87/gin-template/pkg/setting"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

/* Initialize database
- Init a database connection
- Auto migrate all models
*/
func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		setting.CONFIG.DBHost, setting.CONFIG.DBUsername, setting.CONFIG.DBPassword,
		setting.CONFIG.DBName, setting.CONFIG.DBPort, setting.CONFIG.DBTimezone)

	// connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.WithFields(logrus.Fields{
			"type": "connect db error",
		}).Error(err.Error())
		return nil
	}

	DB = db

	return db
}
