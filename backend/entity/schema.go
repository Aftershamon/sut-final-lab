package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name string `valid:"required~Name can not Blank"`
	Email string
	CustomerID string `valid:"matches([L|M|H][0-9])~ผิดรูปแบบ,length(8|8)~ผิดรูปแบบ,matches([L%|M%|H%][0-9])~ผิดรูปแบบ"`
}