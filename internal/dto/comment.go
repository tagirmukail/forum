package dto

import "time"

type CommentRequest struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

type Comments struct {
	Total int       `json:"total"`
	Data  []Comment `json:"data"`
}

type Comment struct {
	ID        string    `json:"id"`
	TopicID   string    `json:"topic_id"`
	UserID    string    `json:"user_id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
