package model

import (
	"time"

	"github.com/tagirmukail/forum/internal/dto"
)

type Comment struct {
	ID        string    `db:"id"`
	TopicID   string    `db:"topic_id"`
	UserID    string    `db:"user_id"`
	Content   string    `db:"content"`
	Author    string    `db:"author"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (c *Comment) ToDTO() dto.Comment {
	if c == nil {
		return dto.Comment{}
	}

	return dto.Comment{
		ID:        c.ID,
		TopicID:   c.TopicID,
		UserID:    c.UserID,
		Author:    c.Author,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
