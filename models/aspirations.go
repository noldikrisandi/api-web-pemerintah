package models

import "time"

type Aspirations struct {
	ID               string    `gorm:"primaryKey" json:"id"`
	Jenis            string    `json:"jenis"`
	Kecamatan        string    `json:"kecamatan"`
	Desa             string    `json:"desa"`
	Keterangan       string    `json:"keterangan"`
	UrlFoto          string    `gorm:"column:url_foto" json:"url_foto"`
	UrlProposal      string    `gorm:"column:url_proposal" json:"url_proposal"`
	Status           string    `json:"status"`
	KeteranganStatus string    `json:"keterangan_status"`
	IdPengirim       string    `gorm:"column:id_pengirim" json:"id_pengirim"` // Sesuaikan nama kolom di DB
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
