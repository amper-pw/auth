package model

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string    `gorm:"size:256;unique;not null"`
	Email    string    `gorm:"size:128;unique;not null"`
	Phone    string    `gorm:"size:128;unique;not null"`
	Password string    `gorm:"size:256;not null"`
	Status   int8      `gorm:"not null"`
}
