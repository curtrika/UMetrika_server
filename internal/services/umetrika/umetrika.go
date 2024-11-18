package umetrika

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/google/uuid"
)

type UMetrika struct {
	log                *slog.Logger
	tp                 testProvider
	schoolInfoProvider schoolInfoProvider
}

type testProvider interface {
	CreateOwner(ctx context.Context, name string, email string, pass_hash []byte) (models.EducationOwner, error)
	CreateTest(ctx context.Context, testName string, description string, testType string, ownerID uuid.UUID) (models.EducationTest, error)
	GetOwner(ctx context.Context, ownerId uuid.UUID) (models.EducationOwner, error)
	GetTestsByOwnerId(ctx context.Context, ownerId uuid.UUID) ([]models.EducationTest, error)
	InsertQuestionsToTest(ctx context.Context, questions []*models.QuestionAnswer) error
	GetFullTestsByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]models.EducationTestFull, error)
}

type schoolInfoProvider interface {
	GetTeacherDisciplinesAndClasses(ctx context.Context, teacherID uuid.UUID) ([]models.TeacherDiscipline, error)
}

func New(
	log *slog.Logger,
	provider testProvider,
	schoolInfoProvider schoolInfoProvider,
) *UMetrika {
	return &UMetrika{
		log:                log,
		tp:                 provider,
		schoolInfoProvider: schoolInfoProvider,
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

func (t *UMetrika) CreateTest(ctx context.Context, testName, description, testType string, ownerId uuid.UUID) (models.EducationTest, error) {
	test, err := t.tp.CreateTest(ctx, testName, description, testType, ownerId)
	if err != nil {
		return models.EducationTest{}, err
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

func (t *UMetrika) InsertQuestionsToTest(ctx context.Context, questions []*models.QuestionAnswer) error {
	err := t.tp.InsertQuestionsToTest(ctx, questions)
	if err != nil {
		return err
	}
	return nil
}

func (t *UMetrika) GetFullTestsByOwnerId(ctx context.Context, ownerId uuid.UUID) ([]models.EducationTestFull, error) {
	tests, err := t.tp.GetFullTestsByOwnerID(ctx, ownerId)
	if err != nil {
		return nil, fmt.Errorf("error while getting full tests by owner id: %w", err)
	}
	return tests, nil
}

func (t *UMetrika) GetTeacherDisciplinesAndClasses(ctx context.Context, teacherID uuid.UUID) ([]models.TeacherDiscipline, error) {
	return t.schoolInfoProvider.GetTeacherDisciplinesAndClasses(ctx, teacherID)
}
