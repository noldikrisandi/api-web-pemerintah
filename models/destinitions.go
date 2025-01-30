package models

type Destinations struct {
	ID               uint   `json:"id" gorm:"primary_key"`
	Judul            string `json:"judul" gorm:"not null"`
	DeskripsiSingkat string `json:"deskripsi_singkat" gorm:"not null"`
	Konten           string `json:"konten" gorm:"not null"`
	Penulis          string `json:"penulis" gorm:"not null"`
	Gambar           string `json:"gambar" gorm:"not null"`
}
