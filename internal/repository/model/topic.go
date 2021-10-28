package model

import (
	"time"

	"github.com/tagirmukail/forum/internal/dto"
)

type Topic struct {
	ID          string    `db:"id"`
	UserID      string    `db:"user_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (t *Topic) ToDTO() dto.Topic {
	if t == nil {
		return dto.Topic{}
	}

	return dto.Topic{
		ID:          t.ID,
		UserID:      t.UserID,
		Name:        t.Name,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

type TopicDetailed struct {
	ID            string    `db:"id"`
	UserID        string    `db:"user_id"`
	Author        string    `db:"author"`
	Name          string    `db:"name"`
	Description   string    `db:"description"`
	TotalComments int       `db:"total_comments"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func (t *TopicDetailed) ToDTO() dto.TopicDetailed {
	if t == nil {
		return dto.TopicDetailed{}
	}

	return dto.TopicDetailed{
		Topic: dto.Topic{
			ID:          t.ID,
			UserID:      t.UserID,
			Name:        t.Name,
			Description: t.Description,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		},
		Author:        t.Author,
		TotalComments: t.TotalComments,
	}
}
