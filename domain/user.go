package domain

type User struct {
	ID              int64        `json:"id"`
	FirstName       string       `json:"first_name"`
	LastName        string       `json:"last_name"`
	Email           string       `json:"email"`
	Password        string       `json:"password"`
	InscriptionDate string       `json:"inscription_date"`
	Subscription    Subscription `json:"subscription"`
	Payments        []Payment    `json:"payments"`
}
