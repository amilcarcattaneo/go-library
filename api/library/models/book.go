package models

import "time"

type Book struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
	Title     string     `json:"title"`
	Author    Author     `json:"author"`
	Thumbnail string     `json:"thumbnail"`
}

func (Book) TableName() string {
	return "books"
}
