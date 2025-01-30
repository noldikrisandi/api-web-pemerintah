package models

import "time"

type Informations struct {
	ID      uint      `json:"id" gorm:"primary_key"`
	Judul   string    `json:"judul" gorm:"type:varchar(255);not null"`
	Konten  string    `json:"konten" gorm:"type:text;not null"`
	Penulis string    `json:"penulis" gorm:"type:varchar(100);not null"`
	Tanggal time.Time `json:"tanggal" gorm:"type:timestamp;default:current_timestamp"`
}
