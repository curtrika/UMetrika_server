package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/curtrika/UMetrika_server/internal/helpers/jwt"
	"github.com/curtrika/UMetrika_server/internal/helpers/logger/sl"
	storage "github.com/curtrika/UMetrika_server/internal/storage/errs"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

// TODO: вынести на кладбище ошибок
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Auth struct {
	log       *slog.Logger
	usrWriter UserWriter
	usrReader UserReader
	appReader AppReader
	tokenTTL  time.Duration
}

type UserWriter interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uuid.UUID, error)
}

type UserReader interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type AppReader interface {
	GetAppById(ctx context.Context, appID int32) (*models.App, error)
}

func New(
	log *slog.Logger,
	userWriter UserWriter,
	userReader UserReader,
	appReader AppReader,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		usrWriter: userWriter,
		usrReader: userReader,
		log:       log,
		appReader: appReader,
		tokenTTL:  tokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, email string, password string, appID int32) (string, error) {
	const op = "Auth.Login"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("attempting to login user")

	// get the user from the database
	user, err := a.usrReader.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", sl.Err(err))

			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		a.log.Error("failed to get user", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	// check the correctness of the received password
	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		a.log.Info("invalid credentials", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	// get information about the application
	app, err := a.appReader.GetAppById(ctx, appID)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user logged in successfully")

	// create an jwt token
	token, err := jwt.NewToken(*user, *app, a.tokenTTL)
	if err != nil {
		a.log.Error("failed to generate token", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, pass string) (uuid.UUID, error) {
	const op = "Auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("registering user")

	// generated hash and salt for password
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost) //TODO: не забыть заменить DefaultCost на что то понадежнее
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.usrWriter.SaveUser(ctx, email, passHash)
	if err != nil {
		log.Error("failed to register user", sl.Err(err))
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
