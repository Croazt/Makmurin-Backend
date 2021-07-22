package dbConfig

import (
	"os"

	"github.com/Croazt/Makmurin/src/utils"
	"github.com/Croazt/Makmurin/src/utils/dbConfig/dbModel"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("DATABASE_HOST")
	name := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	port := os.Getenv("DATABASE_PORT")
	if host == "" {
		host = utils.GoDotEnvVariable("DATABASE_HOST")
		name = utils.GoDotEnvVariable("DATABASE_NAME")
		user = utils.GoDotEnvVariable("DATABASE_USER")
		password = utils.GoDotEnvVariable("DATABASE_PASSWORD")
		port = utils.GoDotEnvVariable("DATABASE_PORT")
	}
	// url := utils.GoDotEnvVariable("DATABSE_URL")

	var err error
	dsn := "host=" + host + " user=" + user + "  password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"

	// dsn := "root:@tcp(127.0.0.1:3306)/ProjectAkhirPemweb?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	if err := DB.AutoMigrate(&dbModel.User{}, &dbModel.Post{}); err != nil {
		panic("Failed to migrate the database")
	}

}
