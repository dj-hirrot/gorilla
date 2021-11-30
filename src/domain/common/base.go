package common

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID  `gorm:"primary_key;" json:"id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	CreatedAt time.Time  `json:"createdAt" example:"2021-11-30T04:17:20.14Z"`
	UpdatedAt time.Time  `json:"updatedAt" example:"2021-11-30T04:17:20.14Z"`
	DeletedAt *time.Time `json:"deletedAt" example:"2021-11-30T04:17:20.14Z"`
}

type BaseExerpt struct {
	ID        uuid.UUID `gorm:"primary_key;" json:"id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	CreatedAt time.Time `json:"createdAt" example:"2021-11-30T04:17:20.14Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2021-11-30T04:17:20.14Z"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.NewV4()
	return
}
