package entity

type Address struct {
	Id          int    `json:"id"`
	UserId      string `json:"user_id"`
	Description string `json:"description"`
	Cep         string `json:"cep"`
	Number      int    `json:"number"`
	Complement  string `json:"complement"`
}
