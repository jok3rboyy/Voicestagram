package repositories

import (
	"Voicestagram/types"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindUserByID retrieves a user by ID from the database
func (repo *UserRepository) FindUserByID(userID uint) (*types.User, error) {
	var user types.User
	result := repo.db.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// FindUserByUsername retrieves a user by username from the database
func (repo *UserRepository) FindUserByUsername(username string) (*types.User, error) {
	var user types.User
	result := repo.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// CreateUser creates a new user in the database
func (repo *UserRepository) CreateUser(user *types.User) error {
	result := repo.db.Create(user)
	return result.Error
}
