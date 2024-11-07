package admin_panel

import (
	"context"
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"log/slog"
)

type AdminPanel struct {
	log      *slog.Logger
	provider Provider
}

type Provider interface {
	//TODO: сюда воообще все методы напишем
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
}

func New(
	log *slog.Logger,
	provider Provider,
) *AdminPanel {
	return &AdminPanel{
		log:      log,
		provider: provider,
	}
}

func (a *AdminPanel) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	newUser, err := a.provider.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
