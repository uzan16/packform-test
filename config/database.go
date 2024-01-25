package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	addr := os.Getenv("DB_ADDR")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	ssl := os.Getenv("DB_SSL")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		addr,
		user,
		pass,
		name,
		port,
		ssl,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
