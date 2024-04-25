package user_entity

import (
	"context"

	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/infra/internal_error"
)

type User struct {
	UserId string
	Name   string
}

type UserRepositoryInterface interface {
	FindUserById(ctx context.Context, userId string) (*User, *internal_error.InternalError)
}
