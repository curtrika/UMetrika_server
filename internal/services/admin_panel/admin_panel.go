package admin_panel

import (
	"log/slog"
)

type AdminPanel struct {
	log      *slog.Logger
	provider Provider
}

type Provider interface{}

func New(
	log *slog.Logger,
	provider Provider,
) *AdminPanel {
	return &AdminPanel{
		log:      log,
		provider: provider,
	}
}

// CreatePsychologicalType creates a new psychological type
// func (a *AdminPanel) CreatePsychologicalType(ctx context.Context, title string) error {
// 	_, err := a.provider.CreatePsychologicalType(ctx, title)
// 	return err
// }
//
// // CreatePsychologicalTest creates a new psychological test
// func (a *AdminPanel) CreatePsychologicalTest(ctx context.Context, title string, typeID int32) error {
// 	_, err := a.provider.CreatePsychologicalTest(ctx, storage.CreatePsychologicalTestParams{
// 		Title:  title,
// 		TypeID: pgtype.Int4{Int32: typeID},
// 	})
// 	return err
// }
//
// // CreateQuestion creates a new question
// func (a *AdminPanel) CreateQuestion(ctx context.Context, questionTitle string, testID int32) error {
// 	_, err := a.provider.CreateQuestion(ctx, storage.CreateQuestionParams{
// 		Title: questionTitle,
// 		// Text:   questionText,
// 		// TestID: testID,
// 	})
// 	return err
// }
//
// // CreateAnswer creates a new answer
// func (a *AdminPanel) CreateAnswer(ctx context.Context, questionID int32, answer string) error {
// 	_, err := a.provider.CreateAnswer(ctx, storage.CreateAnswerParams{
// 		// QuestionID: questionID,
// 		Title: answer,
// 	})
// 	return err
// }
//
// // CreateUser creates a new user
// func (a *AdminPanel) CreateUser(ctx context.Context, username, email string) error {
// 	_, err := a.provider.CreateUser(ctx, storage.CreateUserParams{
// 		// Username: username,
// 		Email: email,
// 	})
// 	return err
// }
//
// // GetApp retrieves an app by its ID
// func (a *AdminPanel) GetApp(ctx context.Context, id int32) (storage.App, error) {
// 	return a.provider.GetApp(ctx, id)
// }
//
// // GetPsychologicalType retrieves a psychological type by its ID
// func (a *AdminPanel) GetPsychologicalType(ctx context.Context, id int32) (storage.PsychologicalType, error) {
// 	return a.provider.GetPsychologicalType(ctx, id)
// }
//
// // GetPsychologicalTest retrieves a psychological test by its ID
// func (a *AdminPanel) GetPsychologicalTest(ctx context.Context, id int32) (storage.PsychologicalTest, error) {
// 	return a.provider.GetPsychologicalTest(ctx, id)
// }
//
// // GetQuestion retrieves a question by its ID
// func (a *AdminPanel) GetQuestion(ctx context.Context, id int32) (storage.Question, error) {
// 	return a.provider.GetQuestion(ctx, id)
// }
//
// // GetUser retrieves a user by their UUID
// func (a *AdminPanel) GetUser(ctx context.Context, id pgtype.UUID) (storage.User, error) {
// 	return a.provider.GetUser(ctx, id)
// }
//
// // ListPsychologicalTypes lists all psychological types
// func (a *AdminPanel) ListPsychologicalTypes(ctx context.Context) ([]storage.PsychologicalType, error) {
// 	return a.provider.ListPsychologicalTypes(ctx)
// }
//
// // ListPsychologicalTests lists all psychological tests
// func (a *AdminPanel) ListPsychologicalTests(ctx context.Context) ([]storage.PsychologicalTest, error) {
// 	return a.provider.ListPsychologicalTests(ctx)
// }
//
// // ListQuestions lists all questions
// func (a *AdminPanel) ListQuestions(ctx context.Context) ([]storage.Question, error) {
// 	return a.provider.ListQuestions(ctx)
// }
//
// // ListAnswers lists all answers
// func (a *AdminPanel) ListAnswers(ctx context.Context) ([]storage.Answer, error) {
// 	return a.provider.ListAnswers(ctx)
// }
//
// // ListUsers lists all users
// func (a *AdminPanel) ListUsers(ctx context.Context) ([]storage.User, error) {
// 	return a.provider.ListUsers(ctx)
// }
//
// // ListApps lists all apps
// func (a *AdminPanel) ListApps(ctx context.Context) ([]storage.App, error) {
// 	return a.provider.ListApps(ctx)
// }
