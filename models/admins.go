package models

type Admins struct {
	ID       string `json:"id" gorm:"primaryKey;type:varchar(255)"`
	Email    string `json:"email" gorm:"unique;not null;type:varchar(255)"`
	Password string `json:"password" gorm:"not null;type:text"`
	Answer   string `json:"answer" gorm:"not null;type:text"`
}
