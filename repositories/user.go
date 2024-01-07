package repositories

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/jok3rboyy/VoiceStagram1/types"
)

// CreateUser maakt een nieuwe gebruiker aan in de database.
func CreateUser(db *gorm.DB, user *types.User, password string) error {
	// Hash de password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Sla de gehashte password op in de user
	user.Password = string(hashedPassword)

	// Sla de user op in de database
	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
