package domain

type Service struct {
	ID           int64        `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Subscription Subscription `json:"subscription"`
}
