package services

import "github.com/SemenVologdin/go-advanced/config"

type Service struct {
	Sender  *Sender
	Hash    *Hash
	Storage *JsonStorage
}

func New(cfg config.App) *Service {
	return &Service{
		Sender:  newSenderService(cfg.Mail),
		Hash:    newHashService(),
		Storage: newJsonStorageService(cfg.Storage.FilePath),
	}
}
