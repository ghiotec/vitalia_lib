package domain

type Subscription struct {
	ID               int64     `json:"id" gorm:"primaryKey"`
	Name             string    `json:"name" gorm:"size:100"`
	Description      string    `json:"description" gorm:"size:100"`
	MonthlyPrice     float64   `json:"monthly_price" gorm:"default:1"`
	DurationInMonths int64     `json:"duration_in_months" gorm:"default:1"`
	Users            []User    `json:"users" gorm:"foreignKey:ID"`
	Services         []Service `json:"services"  gorm:"foreignKey:ID"`
}
