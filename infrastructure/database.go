package infrastructure

import (
	"os"

	"github.com/nezertiam/fiber-erp/internals/core/domain"

	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host='%s' user='%s' password='%s' dbname='%s' port=%s sslmode=disable",
		host, user, password, dbname, port,
	)
	if instance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		db = instance
		log.Println("Database connection established")
	}
}

func GetDB() *gorm.DB {
	return db
}

func MigrateAll() {
	log.Println("Migrating all tables")
	db.AutoMigrate(
		&domain.User{},
	)
}
