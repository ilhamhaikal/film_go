package models

type Cinema struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name" binding:"required" gorm:"not null"`
	Location string  `json:"location" binding:"required" gorm:"not null"`
	Rating   float64 `json:"rating"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
