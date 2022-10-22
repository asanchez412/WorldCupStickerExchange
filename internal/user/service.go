package sticker

import "github.com/asanchez412/WorldCupStickerExchange/pkg/logger"

type Service interface {
}

type service struct {
	repository Repository
	logger     logger.Logger
}

func NewService(r Repository, l logger.Logger) Service {
	return &service{
		repository: r,
		logger:     l,
	}
}
