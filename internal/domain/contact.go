package domain


import "time"

type Contact struct {
ID uint `json:"id" gorm:"primaryKey"`
FirstName string `json:"first_name"`
LastName string `json:"last_name"`
Email string `json:"email" gorm:"index"`
Phone string `json:"phone"`
Company string `json:"company"`
Notes string `json:"notes"`
CreatedAt time.Time `json:"created_at"`
UpdatedAt time.Time `json:"updated_at"`
}