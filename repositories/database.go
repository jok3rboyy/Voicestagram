package repositories

import (
	"fmt"

	"github.com/jok3rboyy/VoiceStagram1/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func DatabaseOpen() {
	dsn := "root:Harvey1994!!@tcp(localhost:3306)/ip_insta_opdracht?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// Print the error and exit the application if the database connection fails
		fmt.Println("Error connecting to the database:", err)
		panic("Failed to connect to the database")
	}

	// Perform database migrations or other setup if needed
	db.AutoMigrate(&types.Post{}, &types.User{})

	// Assign the database connection to the global variable
	Database = db

	// Print a message indicating a successful database connection
	fmt.Println("Connected to the database")

	// Check the connection health
	if err := db.Exec("SELECT 1").Error; err != nil {
		fmt.Println("Error checking the database connection:", err)
		panic("Failed to check the database connection")
	}
	fmt.Println("Database connection is healthy")
}
