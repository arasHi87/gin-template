package setting

import (
	"fmt"

	"github.com/arashi87/gin-template/pkg/model"
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
		CONFIG.DBHost, CONFIG.DBUsername, CONFIG.DBPassword, CONFIG.DBName, CONFIG.DBPort, CONFIG.DBTimezone)

	// connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Init DB error with ", err)
	}

	DB = db

	// migrate models
	DBAutoMigrate()

	return db
}

func DBAutoMigrate() {
	DB.AutoMigrate(&model.UserModel{})
	fmt.Println("Auto migrate all models")
}
