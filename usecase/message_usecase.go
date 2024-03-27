package usecase

import (
	"context"
	"time"

	"github.com/MrBooi/go_chat_backend/domain"
)

type messageUsecase struct {
	messageRepository domain.MessageRepository
	contextTimeout    time.Duration
}

func (mu *messageUsecase) PrivateMessageList(c context.Context, id string) ([]domain.Message, error) {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()

	options := domain.PaginationOptions{
		Page:    20,
		PerPage: 1,
	}

	return mu.messageRepository.PrivateMessageList(ctx, id, options)
}

func (mu *messageUsecase) ConversationMessages(c context.Context, id string, conversationId string) ([]domain.Message, error) {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()
	options := domain.PaginationOptions{
		Page:    20,
		PerPage: 1,
	}
	return mu.messageRepository.ConversationMessages(ctx, id, conversationId, options)
}

func (mu *messageUsecase) Create(c context.Context, body *domain.Message) (domain.Message, error) {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()

	return mu.messageRepository.Create(ctx, body)
}

func NewMessageUsecase(messageRepository domain.MessageRepository, timeout time.Duration) domain.MessageUsecase {
	return &messageUsecase{
		messageRepository: messageRepository,
		contextTimeout:    timeout,
	}
}
