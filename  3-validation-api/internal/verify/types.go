package verify

import (
	"context"
)

type HandlerDeps struct {
	SenderService  Sender
	StorageService StorageService
	HashService    HashService
}

type Sender interface {
	Send(email string, hashEmail string) error
}

type StorageService interface {
	SaveEmailHash(ctx context.Context, email string, hash string) error
	EmailByHash(ctx context.Context, hash string) (string, error)
	DeleteHash(ctx context.Context, hash string) error
}

type HashService interface {
	HashString(text string) string
}
