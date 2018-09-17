package models

import (
	"time"

	"github.com/go-pg/pg/orm"
)

// User define users table structure
type User struct {
	ID        int64
	UserName  string
	Email     string
	Password  string
	Name      string
	Roles     []string
	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time
}

// BeforeInsert hook before insert
func (u *User) BeforeInsert(db orm.DB) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	return nil
}
