package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4() column:id"`
	FirstName string     `json:"first_name" gorm:"column:first_name"`
	LastName  string     `json:"last_name" gorm:"column:last_name"`
	Password  string     `json:"-" gorm:"column:password"`
	Email     string     `json:"email" gorm:"column:email"`
	Birthday  *time.Time `json:"birthday" gorm:"column:birthday"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at"`
}

type UserCreation struct {
	FirstName string     `json:"first_name" form:"first_name" gorm:"column:first_name" binding:"required"`
	LastName  string     `json:"last_name" form:"last_name" gorm:"column:last_name" binding:"required"`
	Password  string     `json:"password" form:"password" gorm:"column:password" binding:"required"`
	Email     string     `json:"email" form:"email" gorm:"column:email" binding:"required"`
	Birthday  *time.Time `json:"birthday" form:"birthday" gorm:"column:birthday"`
}

type TodoItemCreation struct {
	Id          int     `json:"-" gorm:"column:id"`
	Title       string  `json:"title" form:"title" gorm:"column:title;"`
	Description *string `json:"description" form:"description" gorm:"column:description;"`
}

func (User) TableName() string {
	return "users"
}

func (TodoItemCreation) TableName() string {
	return User{}.TableName()
}

func (UserCreation) TableName() string {
	return User{}.TableName()
}
