package user_usecase

import (
	"context"

	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/entity/user_entity"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/infra/internal_error"
)

func NewUserUseCase(userRepository user_entity.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

type UserOutputDto struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
}

type UserUseCaseInterface interface {
	FindUserById(ctx context.Context, id string) (*UserOutputDto, *internal_error.InternalError)
}

func (u *UserUseCase) FindUserById(ctx context.Context, id string) (*UserOutputDto, *internal_error.InternalError) {
	userEntity, err := u.UserRepository.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &UserOutputDto{
		UserId: userEntity.UserId,
		Name:   userEntity.Name,
	}, nil
}
