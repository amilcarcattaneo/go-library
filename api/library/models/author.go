package models

import "time"

type Author struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
	Name      string     `gorm:"column:name" json:"name"`
}

func (Author) TableName() string {
	return "authors"
}
