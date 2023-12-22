package user

import (
	"context"
	"time"

	"github.com/seivanov1986/gocart/internal/repository/user"
	"github.com/seivanov1986/gocart/helpers"
)

type UserCreateIn struct {
	Login     string    `db:"login"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

func (s *service) Create(ctx context.Context, in UserCreateIn) (*int64, error) {
	hash, err := helpers.GenerateHashPassword(in.Password)
	if err != nil {
		return nil, err
	}

	return s.hub.User().Create(ctx, user.UserCreateInput{
		Login:     in.Login,
		Email:     in.Email,
		Password:  hash,
		Active:    in.Active,
		CreatedAt: in.CreatedAt,
	})
}
