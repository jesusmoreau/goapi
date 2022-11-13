package repository

import (
	"context"
	"goapi/internal/entity"
)

const (
	qryInsertUser = `insert into USERS (email, name, password)	values (?, ?, ?);`

	qryGetUserByEmail = `
		select
			id,
			email,
			name,
			password
		from USERS
		where email = ?;`

	qryInsertUserRole = `insert into USER_ROLES (user_id, role_id) values (:user_id, :role_id);`

	qryRemoveUserRole = `delete from USER_ROLES where user_id = :user_id and role_id = :role_id;`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.DB.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.GetContext(ctx, user, qryGetUserByEmail, email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
