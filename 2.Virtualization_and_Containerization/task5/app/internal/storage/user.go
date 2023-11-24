package storage

import (
	"app/internal/db"
	"app/internal/model"
	"context"
)

type UserStorage struct {
	db *db.DB
}

func NewUser(db *db.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

const queryAllSql = "SELECT id, username FROM users OFFSET $1 LIMIT $2"

func (s *UserStorage) QueryAll(ctx context.Context, offset, limit int) (
	[]model.User,
	error,
) {
	rows, err := s.db.Query(ctx, queryAllSql, offset, limit)
	if err != nil {
		return nil, err
	}

	users := make([]model.User, 0)

	for rows.Next() {
		var u model.User
		if err = rows.Scan(&u.ID, &u.Username); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

const insertOneSql = `INSERT INTO users (username) VALUES ($1) RETURNING id;`

func (s *UserStorage) InsertOne(ctx context.Context, username string) (
	*model.User,
	error,
) {
	var u model.User
	u.Username = username

	if err := s.db.QueryRow(
		ctx,
		insertOneSql,
		username,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return &u, nil
}

const deleteOneSql = `DELETE FROM users WHERE id = $1`

func (s *UserStorage) DeleteOne(ctx context.Context, id int) error {
	if _, err := s.db.Exec(ctx, deleteOneSql, id); err != nil {
		return err
	}

	return nil
}
