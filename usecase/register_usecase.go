package usecase

import (
	"context"
	"time"

	"github.com/MrBooi/go_chat_backend/domain"
)

type registerUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func (r *registerUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.userRepository.Create(ctx, user)
}

func (r *registerUsecase) GetUserByUuidOrEmail(c context.Context, uuid string, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.userRepository.GetUserByUuidOrEmail(ctx, uuid, email)
}

func NewRegisterUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RegisterUsecase {
	return &registerUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
