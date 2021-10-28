package logic

import (
	"context"
	"errors"

	"github.com/tagirmukail/forum/internal/repository/model"

	"github.com/tagirmukail/forum/internal/dto"
)

func (l *Logic) validateNewComment(topicID string, comment dto.CommentRequest) error {
	if topicID == "" {
		return errors.New("topic_id is required")
	}

	if comment.UserID == "" {
		return errors.New("user_id is required")
	}

	if comment.Content == "" {
		return errors.New("content is required")
	}

	return nil
}

func (l *Logic) CreateComment(
	ctx context.Context,
	topicID string,
	comment dto.CommentRequest,
) (result dto.Comment, err error) {
	err = l.validateNewComment(topicID, comment)
	if err != nil {
		return result, err
	}

	var commentM model.Comment

	commentM, err = l.d.Repo.NewComment(ctx, model.Comment{
		TopicID: topicID,
		UserID:  comment.UserID,
		Content: comment.Content,
	})
	if err != nil {
		return result, err
	}

	result = commentM.ToDTO()

	return result, nil
}

func (l *Logic) ListComments(ctx context.Context, topicID string, limit, offset int) (result dto.Comments, err error) {
	comments, total, err := l.d.Repo.ListComments(ctx, topicID, limit, offset)
	if err != nil {
		return result, err
	}

	result.Total = total

	for _, comment := range comments {
		result.Data = append(result.Data, comment.ToDTO())
	}

	return result, nil
}
