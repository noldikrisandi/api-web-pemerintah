package models

type Aspirations struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	NIK          string `json:"nik" gorm:"not null"`
	JenisBantuan string `json:"jenis_bantuan" gorm:"not null"`
	Keterangan   string `json:"keterangan" gorm:"not null"`
	LinkGambar   string `json:"link_gambar" gorm:"not null"`
	LinkProposal string `json:"link_proposal" gorm:"not null"`
	Status       string `json:"status" gorm:"default:Pending"`
}
