package entity

import (
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model

	Title string  `json:"title"	valid:"stringlength(3|100)~Title is must be 3 to 100 charaters"`
	Price float64 `json:"price"	valid:"range(50|5000)~Price must be between 50 and 5000"`
	Code  string  `json:"code"	valid:"maxstringlength(8)~Code must start with BK followed by 6 digits (0-9)"`
}
