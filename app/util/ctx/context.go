package ctx

import (
	"context"
	"errors"
)

type contextKey string

const userKey contextKey = "user"

type CtxUser struct {
	ID     int
	Name   string
	ApiKey string
}

func GetCtxUser(ctx context.Context) (*CtxUser, error) {
	v := ctx.Value(userKey)
	user, ok := v.(*CtxUser)
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}
