package entity

type User struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Gender string
	Age    int
}
