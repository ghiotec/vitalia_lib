package domain

type Subscription struct {
	ID               int64     `json:"id" gorm:"primaryKey"`
	Name             string    `json:"name" gorm:"size:100"`
	Description      string    `json:"description" gorm:"size:100"`
	MonthlyPrice     float64   `json:"monthly_price" gorm:"default:1"`
	DurationInMonths int64     `json:"duration_in_months" gorm:"default:1"`
	Services         []Service `json:"services" gorm:"many2many:subscription_services;"` // N -- N
}
