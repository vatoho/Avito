package dto

import "github.com/ilyushkaaa/banner-service/internal/user/model"

type UserDB struct {
	Token string `db:"token"`
	Role  string `db:"role"`
	TagID uint64 `db:"tag_id"`
}

func ConvertToUser(userDB UserDB) model.User {
	return model.User{
		Token: userDB.Token,
		Role:  userDB.Role,
		TagID: userDB.TagID,
	}
}
