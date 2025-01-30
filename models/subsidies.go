package models

type Subsidies struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	NIK          string `json:"nik" gorm:"unique;not null"`
	JenisBantuan string `json:"jenis_bantuan" gorm:"not null"`
	Keterangan   string `json:"keterangan" gorm:"not null"`
	LinkProposal string `json:"link_proposal" gorm:"not null"`
	Status       string `json:"status" gorm:"default:Pending"`
}
