package models

import "time"

type UserBook struct {
	UserID    uint64     `json:"user_id"`
	BookID    uint64     `json:"book_id"`
	Available bool       `json:"available"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func (UserBook) TableName() string {
	return "user_books"
}
