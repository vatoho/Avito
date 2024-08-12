package storage

import (
	"context"
	"errors"

	"github.com/ilyushkaaa/banner-service/internal/user/model"
	"github.com/ilyushkaaa/banner-service/internal/user/storage"
	"github.com/ilyushkaaa/banner-service/internal/user/storage/database/dto"
	"github.com/jackc/pgx/v4"
)

func (s *UserStorageDB) GetUserByToken(ctx context.Context, token string) (*model.User, error) {
	var userDB dto.UserDB
	err := s.db.Get(ctx, &userDB,
		`SELECT token, tag_id, role 
                FROM users WHERE token = $1`, token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrUserNotFound
		}
		return nil, err
	}
	user := dto.ConvertToUser(userDB)
	return &user, nil
}
