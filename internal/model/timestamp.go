package model

import (
	"time"

	"gorm.io/gorm"
)

type Timestamp struct {
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

func (m *Timestamp) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now().UTC()
	m.UpdatedAt = m.CreatedAt
	return nil
}

func (m *Timestamp) BeforeSave(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	m.UpdatedAt = now
	return nil
}

func (m *Timestamp) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now().UTC()
	return nil
}
