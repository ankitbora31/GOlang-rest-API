package entity

type Address struct {
	ID     int `gorm:"primaryKey"`
	City   string
	Pin    string
	State  string
	UserID string `gorm:"Column:user_id"`
}
