package domain

type Subscription struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	MonthlyPrice     float64   `json:"monthly_price"`
	DurationInMonths int64     `json:"duration_in_months"`
	Users            []User    `json:"users"`
	Services         []Service `json:"services"`
}
