package models

import (
	"github.com/google/uuid"
)

type Users struct {
	ID       uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	NIK      string    `json:"nik"`
	NOKK     string    `json:"nokk"`
	WA       string    `json:"wa"`
	NAMA     string    `json:"nama"`
}
