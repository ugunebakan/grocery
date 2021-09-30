package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
  u2, _ := uuid.NewV4()
  base.ID = u2
  return
}

