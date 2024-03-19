package usecase

import (
	"context"
	"time"

	"github.com/MrBooi/go_chat_backend/domain"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func (uu *userUsecase) UpdateUser(c context.Context, id string, body domain.UpdateUserRequest) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.UpdateUser(ctx, id, body)
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}

}
