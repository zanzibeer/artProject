package user

import (
	"artProject/pkg/logging"
	"context"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	// TODO dor next one
	return
}
