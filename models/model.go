package models

type Registration struct {
	id       string
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tweet struct {
	id          string
	Title       string `json:"title"`
	Discription string `json:"discription"`
}
