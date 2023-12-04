package models

import (
	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	Employee_id string `json:"employee_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Birthday    string `json:"birthday" validate:"required"`
	Age         int    `json:"age" validate:"required,min=1"`
	Email       string `json:"email" validate:"required"`
	Tel         string `json:"tel" validate:"required,min=10,max=10"`
}

type UsersRes struct {
	Employee_id string `json:"employee_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Birthday    string `json:"birthday" validate:"required"`
	Age         int    `json:"age" validate:"required,min=1"`
	Email       string `json:"email" validate:"required"`
	Tel         string `json:"tel" validate:"required,min=10,max=10"`
	Generation  string `json:"generation" validate:"required"`
}

type ResultData struct {
	All_users    int        `json:"count"`
	Data         []UsersRes `json:"data"`
	GenZ         int        `json:"GenZ"`
	GenY         int        `json:"GenY"`
	GenX         int        `json:"GenX"`
	BabyBoomer   int        `json:"BabyBoomer"`
	GIGeneration int        `json:"GIGeneration"`
}
