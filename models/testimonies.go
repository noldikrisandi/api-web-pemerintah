package models

type Testimonies struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Nama   string `json:"nama" gorm:"not null"`
	Kota   string `json:"kota" gorm:"not null"`
	Konten string `json:"konten" gorm:"not null"`
	Gambar string `json:"gambar" gorm:"not null"`
}
