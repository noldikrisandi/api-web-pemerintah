package models

type Admincontroller struct {
	ID         string `json:"id" gorm:"column:id"`
	IDAdmin    string `json:"idadmin" gorm:"column:idadmin"`
	IDAspirasi string `json:"idaspirasi" gorm:"column:idaspirasi"`
}

// menyesuaikan dengan nama tabel di database
func (Admincontroller) TableName() string {
	return "admincontroller"
}
