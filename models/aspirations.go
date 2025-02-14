package models

type Aspirations struct {
	ID           string `json:"id" gorm:"primary_key"`
	NIK          string `json:"nik"`
	JenisBantuan string `json:"jenis_bantuan"`
	Keterangan   string `json:"keterangan"`
	LinkGambar   string `json:"link_gambar"`
	LinkProposal string `json:"link_proposal"`
	Status       string `json:"status" gorm:"default:Pending"`
}
