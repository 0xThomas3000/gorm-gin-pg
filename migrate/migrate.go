package main

import (
	"fmt"
	"log"

	"github.com/0xThomas3000/gorm-gin-pg/initializers"
	"github.com/0xThomas3000/gorm-gin-pg/models"
)

// 1. Load the environment variables and create a connection pool to the Postgres database
func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

// The AutoMigrate():
//   - Create tables, missing foreign keys, constraints, columns, and indexes.
//   - Update the existing column’s type if its size, precision, and nullable were changed.
func main() {
	// 2. evoked the AutoMigrate() to create the DB migration, then push the changes to the database.
	// 	=> "users" table will have columns added by the GORM migration tool.
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("👍 Migration complete")
}
