package models

type Packages struct {
	ID               uint    `json:"id" gorm:"primary_key"`
	Nama             string  `json:"nama" gorm:"not null"`
	DeskripsiSingkat string  `json:"deskripsi_singkat" gorm:"not null"`
	Konten           string  `json:"konten" gorm:"not null"`
	Gambar           string  `json:"gambar" gorm:"not null"`
	Biaya            float64 `json:"biaya" gorm:"not null"`
}
