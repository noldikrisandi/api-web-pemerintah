package models

import "time"

type Users struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	NIK       string    `json:"nik"`
	NOKK      string    `json:"nokk"`
	WA        string    `json:"wa"`
	Nama      string    `json:"nama"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
