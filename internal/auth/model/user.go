package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FullName  string    `gorm:"size:100;not null" json:"full_name"`
	Email     string    `gorm:"size:100;unique;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	RoleID    uint      `gorm:"not null" json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
