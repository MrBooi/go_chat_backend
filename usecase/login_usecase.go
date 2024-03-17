package usecase

import (
	"context"
	"time"

	"github.com/MrBooi/go_chat_backend/domain"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}


func (lu *loginUsecase) GetUserByUuid(c context.Context, uuid string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetUserByUuid(ctx, uuid)
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
