package repositories

import (
	"github.com/jok3rboyy/VoiceStagram1/types"
	"gorm.io/gorm"
)

type VoiceMessageRepository struct {
	db *gorm.DB
}

// NewVoiceMessageRepository maakt een nieuwe VoiceMessageRepository aan.
func NewVoiceMessageRepository(db *gorm.DB) *VoiceMessageRepository {
	return &VoiceMessageRepository{db: db}
}

// CreateVoiceMessage maakt een nieuwe voicemessage aan in de database.
func (r *VoiceMessageRepository) CreateVoiceMessage(voiceMessage *types.Post) error {
	return r.db.Create(voiceMessage).Error
}

// GetVoiceMessagesByUsername pakt alle voicemessages van een bepaalde user
func (r *VoiceMessageRepository) GetVoiceMessagesByUsername(username string) ([]types.Post, error) {
	var voiceMessages []types.Post
	err := r.db.Where("username = ?", username).Find(&voiceMessages).Error
	return voiceMessages, err
}
