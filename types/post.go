package types

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Username             string
	VoiceMessage         string
	VoiceMessageFilePath string
}
