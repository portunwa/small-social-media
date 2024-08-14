package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content string `gorm:"not null;default:null"`
	UserID  uint
	User 	User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
}

type PostResponse struct {
    ID          uint        `json:"id"`
    Content     string      `json:"content"`
    CreatedAt   time.Time   `json:"createdAt"`
    UpdatedAt   time.Time   `json:"updatedAt"`
    User        UserResponse `json:"user"`
}

func (p Post) ToResponse() PostResponse {
    return PostResponse{
        ID:          p.ID,
        Content:     p.Content,
        CreatedAt:   p.CreatedAt,
        UpdatedAt:   p.UpdatedAt,
        User:        p.User.ToResponse(),
    }
}