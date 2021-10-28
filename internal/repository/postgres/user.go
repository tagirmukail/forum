package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/tagirmukail/forum/internal/repository/model"
)

type UserPostgres struct {
	db             *sqlx.DB
	stmtInsertUser *sqlx.NamedStmt
}

const queryInsertUser = `INSERT INTO users(username) VALUES (:username) 
RETURNING *`

func (u *UserPostgres) NewUser(ctx context.Context, user model.User) (model.User, error) {
	var result model.User

	err := u.stmtInsertUser.GetContext(ctx, &result, user)

	return result, err
}

func (u *UserPostgres) prepareStmt() (err error) {
	u.stmtInsertUser, err = u.db.PrepareNamed(queryInsertUser)
	if err != nil {
		return err
	}

	return nil
}
