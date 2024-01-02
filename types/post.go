package types

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Username uint   // Verwijzing naar de gebruiker die het bericht heeft geplaatst
	Message  string `json:"message"`
}
