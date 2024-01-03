// repositories/user.go

package repositories

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/jok3rboyy/VoiceStagram1/types"
)

// CreateUser creates a new user in the database
func CreateUser(db *gorm.DB, user *types.User, password string) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Set the hashed password in the user object
	user.Password = string(hashedPassword)

	// Save the user to the database
	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
