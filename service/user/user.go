package user

import (
	"time"
)

type User struct {
	ID             int64     `json:"idUser"`
	Name           string    `json:"name" bindinf:"required"`
	Username       string    `json:"username"`
	Email          string    `json:"email" binding:"required"`
	Password       string    `json:"password" binding:"required"`
	Phone          string    `json:"phone" binding:"required"`
	Birthday       time.Time `json:"birthday"`
	PhoneVerified  bool      `json:"phoneVerified"`
	ProfilePicture string    `json:"profilePicture"`
	Gender         string    `json:"gender"`
	CreateAt       time.Time `json:"createAt"`
	UpdateAt       time.Time `json:"updateAt"`
	IsDelete       bool      `json:"isDelete"`
}

func (s *Service) NewUser() *User {

	return &User{
		CreateAt:      time.Now(),
		IsDelete:      false,
		PhoneVerified: false,
	}
}
