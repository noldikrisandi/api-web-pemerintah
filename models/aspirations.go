package models

import "time"

type Aspirations struct {
	ID               string    `gorm:"primaryKey" json:"id"`
	Jenis            string    `json:"jenis"`
	Kecamatan        string    `json:"kecamatan"`
	Desa             string    `json:"Desa"`
	Keterangan       string    `json:"keterangan"`
	UrlFoto          string    `gorm:"column:url_foto" json:"url_foto"`         // Sesuaikan nama kolom di DB
	UrlProposal      string    `gorm:"column:url_proposal" json:"url_proposal"` // Sesuaikan nama kolom di DB
	Status           string    `json:"status"`
	KeteranganStatus string    `json:"keterangan_status"`
	IdPengirim       string    `json:"id_pengirim"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
