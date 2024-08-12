package delivery

import (
	"context"
	"fmt"

	"github.com/ilyushkaaa/banner-service/internal/user/model"
)

type KeyUser string

const UserKey KeyUser = "user"

func getUserFromContext(ctx context.Context) (*model.User, error) {
	user, ok := ctx.Value(UserKey).(*model.User)
	if !ok {
		return nil, fmt.Errorf("no user in context")
	}
	return user, nil
}
