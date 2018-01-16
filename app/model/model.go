package model

import "time"

// Model is the base model type
type Model struct {
	ID        uint       `json:"id" gorm:"primary_key" sql:"name:id"`
	CreatedAt time.Time  `json:"createdAt" sql:"name:createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" sql:"name:updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"name:deletedAt"`
}
