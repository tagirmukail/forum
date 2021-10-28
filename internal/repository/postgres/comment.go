package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/tagirmukail/forum/internal/repository/model"
)

type CommentPostgres struct {
	db                *sqlx.DB
	insertCommentStmt *sqlx.NamedStmt
	totalCommentsStmt *sqlx.Stmt
	listCommentsStmt  *sqlx.Stmt
}

const queryInsertComment = `INSERT INTO comments(topic_id, user_id, content)
VALUES (:topic_id, :user_id, :content) RETURNING *`

func (c *CommentPostgres) NewComment(ctx context.Context, comment model.Comment) (model.Comment, error) {
	var result model.Comment
	err := c.insertCommentStmt.GetContext(ctx, &result, comment)

	return result, err
}

const (
	queryTotalComments = `SELECT COUNT(*) FROM comments WHERE topic_id=$1`
	queryListComments  = `SELECT c.id,
       c.topic_id,
       c.user_id,
       c.content,
       c.created_at,
       c.updated_at,
       u.username as author
FROM comments as c
         JOIN users u on u.id = c.user_id
WHERE topic_id = $1
LIMIT $2 OFFSET $3`
)

func (c *CommentPostgres) ListComments(
	ctx context.Context,
	topicID string,
	limit, offset int,
) (result []model.Comment, total int, err error) {
	err = c.totalCommentsStmt.GetContext(ctx, &total, topicID)
	if err != nil {
		return nil, 0, err
	}

	if limit < 1 {
		limit = defaultLimitPerPage
	}

	err = c.listCommentsStmt.SelectContext(ctx, &result, topicID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

func (c *CommentPostgres) prepareStmt() (err error) {
	c.insertCommentStmt, err = c.db.PrepareNamed(queryInsertComment)
	if err != nil {
		return err
	}

	c.totalCommentsStmt, err = c.db.Preparex(queryTotalComments)
	if err != nil {
		return err
	}

	c.listCommentsStmt, err = c.db.Preparex(queryListComments)
	if err != nil {
		return err
	}

	return nil
}
