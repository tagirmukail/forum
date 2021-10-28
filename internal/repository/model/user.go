package model

import (
	"time"

	"github.com/tagirmukail/forum/internal/dto"
)

type User struct {
	ID        string    `db:"id"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u *User) ToDTO() dto.User {
	if u == nil {
		return dto.User{}
	}

	return dto.User{
		ID:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
