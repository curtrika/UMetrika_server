package testsback

import (
	"fmt"
	"log/slog"

	"github.com/curtrika/UMetrika_server/internal/domain/models"
)

/*
TODO
1) Получение списка тестов по id психолога
2) Добавление нового теста от психолога
3) Редактирование теста от психолога
4) Получение содержимого теста по id теста
5) Получение всех классов школы
6) Получение списка тестов и классов, для которых он проводился
7) Получение результатов теста по каждому ученику из конкретного класса
8) Получение результатов выбора по каждому вопросу из теста в конкретном классе
*/

type TestService struct {
	log           *slog.Logger
	testsProvider testProvider
}

type testProvider interface {
	GetTestsByOwnerId(ownerId int) ([]models.PsychologicalTest, error)
}

func NewTestService(tp testProvider) TestService {
	return TestService{testsProvider: tp}
}

func (t *TestService) GetTestsByOwnerId(ownerId int) ([]models.PsychologicalTest, error) {
	tests, err := t.testsProvider.GetTestsByOwnerId(ownerId)
	if err != nil {
		return nil, fmt.Errorf("could not get tests by owner id: %w", err)
	}
	return tests, nil
}
