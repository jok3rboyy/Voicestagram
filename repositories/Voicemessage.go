package repositories

import (
	"github.com/jok3rboyy/VoiceStagram1/types"
	"gorm.io/gorm"
)

// VoiceMessageRepository represents a repository for VoiceMessage operations
type VoiceMessageRepository struct {
	db *gorm.DB
}

// NewVoiceMessageRepository creates a new VoiceMessageRepository
func NewVoiceMessageRepository(db *gorm.DB) *VoiceMessageRepository {
	return &VoiceMessageRepository{db: db}
}

// CreateVoiceMessage creates a new voice message in the database
func (r *VoiceMessageRepository) CreateVoiceMessage(voiceMessage *types.Post) error {
	return r.db.Create(voiceMessage).Error
}

// GetVoiceMessagesByUserID returns all voice messages for a specific user from the database
func (r *VoiceMessageRepository) GetVoiceMessagesByUserID(userID uint) ([]types.Post, error) {
	var voiceMessages []types.Post
	err := r.db.Where("user_id = ?", userID).Find(&voiceMessages).Error
	return voiceMessages, err
}
