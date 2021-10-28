package repository

import (
	"context"

	"github.com/tagirmukail/forum/internal/repository/model"
)

type Repository interface {
	UserRepository
	TopicRepository
	CommentRepository
}

type UserRepository interface {
	NewUser(ctx context.Context, user model.User) (model.User, error)
}

type TopicRepository interface {
	NewTopic(ctx context.Context, topic model.Topic) (model.Topic, error)
	ListTopics(ctx context.Context, limit, offset int) ([]model.Topic, int, error)
	GetTopic(ctx context.Context, id string) (model.TopicDetailed, error)
}

type CommentRepository interface {
	NewComment(ctx context.Context, comment model.Comment) (model.Comment, error)
	ListComments(ctx context.Context, topicID string, limit, offset int) ([]model.Comment, int, error)
}
