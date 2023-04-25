package domain

type Status string

const (
	Approve Status = "approve"
	Decline Status = "decline"
	Pending Status = "pending"
)

type Payment struct {
	ID     int64   `json:"id"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
	State  Status  `json:"status"`
	User   User    `json:"user"`
}
