package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name string `valid:"required~Name can not Blank"`
	Email string
	CustomerID string 
}