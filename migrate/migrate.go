package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/uzan16/packform-test/config"
	"github.com/uzan16/packform-test/models"
	"gorm.io/gorm"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	db, err := config.ConnectToDB()
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}
	// migrate structure
	if err := db.AutoMigrate(
		&models.CustomerCompanies{},
		&models.Customers{},
		&models.Orders{},
		&models.OrderItems{},
		&models.Deliveries{},
	); err != nil {
		log.Fatal("Failed to migrate: ", err)
	}

	// seed data
	seed(db, "customer_companies", []models.CustomerCompanies{})
	seed(db, "customers", []models.Customers{})
	seed(db, "orders", []models.Orders{})
	seed(db, "order_items", []models.OrderItems{})
	seed(db, "deliveries", []models.Deliveries{})
}

func seed[T any](db *gorm.DB, fileName string, model []T) {
	file, err := os.Open(fmt.Sprintf("./migrate/seed/%s.csv", fileName))
	if err != nil {
		log.Fatal("Failed to load file: ", err)
	}
	defer file.Close()

	err = gocsv.Unmarshal(file, &model)
	if err != nil {
		log.Fatal("Failed to load csv: ", err)
	}

	result := db.Create(model)
	if result.Error != nil {
		log.Fatal("Failed to insert: ", err)
	}
}
