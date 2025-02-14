package models

type Users struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	NIK      string `json:"nik"`
	NOKK     string `json:"nokk"`
	WA       string `json:"wa"`
	Nama     string `json:"nama"`
}
