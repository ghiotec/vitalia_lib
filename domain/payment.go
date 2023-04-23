package domain

type State string

const (
	Approve State = "approve"
	Decline State = "decline"
	Pending State = "pending"
)

type Payment struct {
	ID     int64   `json:"id"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
	State  State   `json:"state"`
	User   User    `json:"user"`
}
