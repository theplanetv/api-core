package models

import (
	"github.com/google/uuid"
)

type JourneyTag struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name string    `gorm:"type:text;unique;not null"`
}
