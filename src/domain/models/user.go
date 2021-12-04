package models

import (
	"github.com/dj-hirrot/gorilla/src/domain/common"
)

type UserAttributes struct {
	Name string `gorm:"column:name" json:"name" example:"Hiroto Shibutani"`
}

type User struct {
	common.Base
	UserAttributes
}

type UserExcerpt struct {
	common.BaseExerpt
	UserAttributes
}

type Users []User

// If you want to use the parent class by inheriting from it, use the following description.
// type Profile struct {
// 	common.Base
// 	Name   string    `gorm:"column:name;size:128;not null;"`
// 	UserID uuid.UUID `gorm:"type:uuid;column:user_foreign_key;not null;"`
// }
