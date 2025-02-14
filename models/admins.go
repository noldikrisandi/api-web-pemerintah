package models

type Admins struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Answer   string `json:"answer"`
}
