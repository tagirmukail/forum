package dto

import "time"

type TopicRequest struct {
	UserID      string `json:"user_id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Topics struct {
	Total int     `json:"total"`
	Data  []Topic `json:"data"`
}

type TopicDetailed struct {
	Topic
	Author        string `json:"author"`
	TotalComments int    `json:"total_comments"`
}

type Topic struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
