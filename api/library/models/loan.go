package models

import "time"

type UserLoan struct {
	UserID       uint64     `json:"user_id"`
	BookID       uint64     `json:"book_id"`
	DueDate      time.Time  `json:"due_date"`
	UserIDLender uint64     `json:"user_id_lender"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func (UserLoan) TableName() string {
	return "user_loans"
}
