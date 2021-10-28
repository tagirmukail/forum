package logic

import (
	"context"
	"errors"

	"github.com/tagirmukail/forum/internal/dto"
	"github.com/tagirmukail/forum/internal/repository/model"
)

func (l *Logic) validateNewTopic(topic dto.TopicRequest) error {
	if topic.Name == "" {
		return errors.New("name is required")
	}

	if topic.Description == "" {
		return errors.New("description is required")
	}

	if topic.UserID == "" {
		return errors.New("user_id is required")
	}

	return nil
}

func (l *Logic) CreateTopic(ctx context.Context, topic dto.TopicRequest) (result dto.Topic, err error) {
	err = l.validateNewTopic(topic)
	if err != nil {
		return result, err
	}

	var topicM model.Topic

	topicM, err = l.d.Repo.NewTopic(ctx, model.Topic{
		UserID:      topic.UserID,
		Name:        topic.Name,
		Description: topic.Description,
	})
	if err != nil {
		return result, err
	}

	result = topicM.ToDTO()

	return result, nil
}

func (l *Logic) ListTopics(ctx context.Context, limit, offset int) (result dto.Topics, err error) {
	topics, total, err := l.d.Repo.ListTopics(ctx, limit, offset)
	if err != nil {
		return result, err
	}

	result.Total = total

	result.Data = make([]dto.Topic, 0, len(topics))
	for _, topic := range topics {
		result.Data = append(result.Data, topic.ToDTO())
	}

	return result, nil
}

func (l *Logic) GetTopic(ctx context.Context, topicID string) (result dto.TopicDetailed, err error) {
	var topicM model.TopicDetailed

	topicM, err = l.d.Repo.GetTopic(ctx, topicID)
	if err != nil {
		return result, err
	}

	result = topicM.ToDTO()

	return result, nil
}
