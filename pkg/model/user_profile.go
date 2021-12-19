package model

import (
	"github.com/google/uuid"
	"time"
)

type UserProfile struct {
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserId     uuid.UUID
	User       User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FirstName  string `gorm:"size:256"`
	LastName   string `gorm:"size:256"`
	MiddleName string `gorm:"size:256"`
	CreatedAt  time.Time
	UpdatedAt  int
}
