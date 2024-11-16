package umetrika

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/google/uuid"
)

type UMetrika struct {
	log *slog.Logger
	tp  testProvider
}

type testProvider interface {
	CreateOwner(ctx context.Context, name string, email string, pass_hash []byte) (models.EducationOwner, error)
	CreateTest(ctx context.Context, testName string, description string, testType string) (models.EducationTest, error)
	GetOwner(ctx context.Context, ownerId uuid.UUID) (models.EducationOwner, error)
	GetTestsByOwnerId(ctx context.Context, ownerId uuid.UUID) ([]models.EducationTest, error)
}

func New(
	log *slog.Logger,
	provider testProvider,
) *UMetrika {
	return &UMetrika{
		log: log,
		tp:  provider,
	}
}

func (t *UMetrika) CreateOwner(ctx context.Context, name string, email string, pass_hash []byte) (models.EducationOwner, error) {
	owner, err := t.tp.CreateOwner(ctx, name, email, pass_hash)
	if err != nil {
		return models.EducationOwner{}, err
	}
	return owner, nil
}

func (t *UMetrika) GetOwner(ctx context.Context, ownerId uuid.UUID) (models.EducationOwner, error) {
	owner, err := t.tp.GetOwner(ctx, ownerId)
	if err != nil {
		return models.EducationOwner{}, err
	}
	return owner, nil
}

func (t *UMetrika) CreateTest(ctx context.Context, testName, description, testType string) (models.EducationTest, error) {
	test, err := t.tp.CreateTest(ctx, testName, description, testType)
	if err != nil {
		return models.EducationTest{}, nil
	}
	return test, nil
}

func (t *UMetrika) GetTestsByOwnerId(ctx context.Context, ownerId uuid.UUID) ([]models.EducationTest, error) {
	tests, err := t.tp.GetTestsByOwnerId(ctx, ownerId)
	if err != nil {
		return nil, fmt.Errorf("could not get tests by owner id: %w", err)
	}
	return tests, nil
}
