package model

type User struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
}

type UserResponse struct {
	Username uint   `json:"username" gorm:"primaryKey"`
	Password string `json:"password" gorm:"unique"`
}
