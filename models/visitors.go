package models

type Visitors struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	Nama      string `json:"nama" gorm:"not null"`
	Kota      string `json:"kota" gorm:"not null"`
	Pekerjaan string `json:"pekerjaan" gorm:"not null"`
	WA        string `json:"wa" gorm:"not null"`
}
