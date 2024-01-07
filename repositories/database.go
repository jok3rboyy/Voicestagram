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
		// Print error en stop de applicatie als ie niet werkt
		fmt.Println("Error connecting to the database:", err)
		panic("Failed to connect to the database")
	}

	// migreer database
	db.AutoMigrate(&types.Post{}, &types.User{})

	// zorg dat de database een eigen variabele krijgt
	Database = db

	// Print dat hij het doet
	fmt.Println("Connected to database")

	// Check of er iets geselecteerd kan worden uit de database
	if err := db.Exec("SELECT 1").Error; err != nil {
		fmt.Println("Error checking the database connection:", err)
		panic("Failed to check the database connection")
	}
	fmt.Println("Database connection is working")
}
