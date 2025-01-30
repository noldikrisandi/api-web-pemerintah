package models

type Users struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
	NIK      string `json:"nik"`
	NOKK     string `json:"nokk"`
	WA       string `json:"wa"`
}