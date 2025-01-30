package models

type Orders struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	IDPackage  uint   `json:"id_package"`
	IDVisitor  uint   `json:"id_visitor"`
	Biaya      string `json:"biaya"`
	Pembayaran string `json:"pembayaran"`
	WA         string `json:"wa"`
	Status     string `json:"status"`
}
