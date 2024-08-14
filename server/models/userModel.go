package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null;default:null"`
	Password string `gorm:"not null;default:null"`
	DisplayName string `gorm:"not null;default:null"`
}

type UserResponse struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
}

func (u User) ToResponse() UserResponse {
    return UserResponse{
        ID:          u.ID,
        Username:    u.Username,
        DisplayName: u.DisplayName,
    }
}