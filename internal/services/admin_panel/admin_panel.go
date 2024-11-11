package admin_panel

import (
	"context"
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/google/uuid"
	"log/slog"
)

type AdminPanel struct {
	log      *slog.Logger
	provider Provider
}

type Provider interface {
	//TODO: сюда временно воообще все методы напишем
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	ReadUser(ctx context.Context, userID uuid.UUID) (*models.User, error)
	UpdateUser(ctx context.Context, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
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

func (a *AdminPanel) ReadUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	user, err := a.provider.ReadUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *AdminPanel) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	updatedUser, err := a.provider.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (a *AdminPanel) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	err := a.provider.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
