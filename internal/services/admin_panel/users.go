package admin_panel

import (
	"context"
	"github.com/curtrika/UMetrika_server/internal/domain/models"
)

type UserProvider interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
}
