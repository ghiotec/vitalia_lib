package domain

type AccessHistory struct {
	ID      int64   `json:"id"`
	Date    string  `json:"date"`
	User    User    `json:"user"`
	Service Service `json:"service"`
}
