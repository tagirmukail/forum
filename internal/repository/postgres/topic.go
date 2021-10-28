package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/tagirmukail/forum/internal/repository/model"
)

type TopicPostgres struct {
	db              *sqlx.DB
	stmtInsertTopic *sqlx.NamedStmt
	stmtTotalTopics *sqlx.Stmt
	stmtListTopics  *sqlx.Stmt
	stmtGetTopic    *sqlx.Stmt
}

const queryInsertTopic = `INSERT INTO topics(user_id, name, description)
VALUES(:user_id, :name, :description) RETURNING *`

func (t *TopicPostgres) NewTopic(ctx context.Context, topic model.Topic) (model.Topic, error) {
	var result model.Topic

	err := t.stmtInsertTopic.GetContext(ctx, &result, topic)

	return result, err
}

const (
	queryTotalTopics = `SELECT COUNT(*) FROM topics`
	queryListTopics  = `SELECT * FROM topics LIMIT $1 OFFSET $2`
)

func (t *TopicPostgres) ListTopics(ctx context.Context, limit, offset int) ([]model.Topic, int, error) {
	var total int

	if limit < 1 {
		limit = defaultLimitPerPage
	}

	err := t.stmtTotalTopics.GetContext(ctx, &total)
	if err != nil {
		return nil, 0, err
	}

	var result []model.Topic

	err = t.stmtListTopics.SelectContext(ctx, &result, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

const queryGetTopic = `SELECT t.id,
       t.user_id,
       u.username  as author,
       t.name,
       t.description,
       t.created_at,
       t.updated_at,
       COUNT(c.*) as total_comments
FROM topics as t
         LEFT JOIN users as u on t.user_id = u.id
         LEFT JOIN comments as c on t.id = c.topic_id
WHERE t.id = $1
GROUP BY t.id, u.username`

func (t *TopicPostgres) GetTopic(ctx context.Context, id string) (model.TopicDetailed, error) {
	var result model.TopicDetailed

	err := t.stmtGetTopic.GetContext(ctx, &result, id)

	return result, err
}

func (t *TopicPostgres) prepareStmt() (err error) {
	t.stmtInsertTopic, err = t.db.PrepareNamed(queryInsertTopic)
	if err != nil {
		return err
	}

	t.stmtTotalTopics, err = t.db.Preparex(queryTotalTopics)
	if err != nil {
		return err
	}

	t.stmtListTopics, err = t.db.Preparex(queryListTopics)
	if err != nil {
		return err
	}

	t.stmtGetTopic, err = t.db.Preparex(queryGetTopic)
	if err != nil {
		return err
	}

	return nil
}
