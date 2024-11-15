// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import (
	models "github.com/curtrika/UMetrika_server/internal/domain/models"
	postgres "github.com/curtrika/UMetrika_server/internal/repository/postgres"
	sqlc "github.com/curtrika/UMetrika_server/internal/repository/postgres/sqlc"
)

type ConverterImpl struct{}

func (c *ConverterImpl) AnswerDBToModel(source sqlc.Answer) models.Answer {
	var modelsAnswer models.Answer
	modelsAnswer.ID = source.ID
	modelsAnswer.NextAnswerID = postgres.Int4ToInt(source.NextAnswerID)
	modelsAnswer.Title = source.Title
	return modelsAnswer
}

func (c *ConverterImpl) AnswerModelToDB(source models.Answer) sqlc.Answer {
	var postgresAnswer sqlc.Answer
	postgresAnswer.ID = source.ID
	postgresAnswer.NextAnswerID = postgres.IntToInt4(source.NextAnswerID)
	postgresAnswer.Title = source.Title
	return postgresAnswer
}

func (c *ConverterImpl) PsychologicalPerfomanceDBToModel(source sqlc.PsychologicalPerformance) models.PsychologicalPerformance {
	var modelsPsychologicalPerformance models.PsychologicalPerformance
	modelsPsychologicalPerformance.ID = source.ID
	modelsPsychologicalPerformance.OwnerID = source.OwnerID
	modelsPsychologicalPerformance.PsychologicalTestID = postgres.Int4ToInt(source.PsychologicalTestID)
	modelsPsychologicalPerformance.StartedAt = postgres.TimestamptzToTime(source.StartedAt)
	return modelsPsychologicalPerformance
}

func (c *ConverterImpl) PsychologicalPerfomanceModelToDB(source models.PsychologicalPerformance) sqlc.PsychologicalPerformance {
	var postgresPsychologicalPerformance sqlc.PsychologicalPerformance
	postgresPsychologicalPerformance.ID = source.ID
	postgresPsychologicalPerformance.OwnerID = source.OwnerID
	postgresPsychologicalPerformance.PsychologicalTestID = postgres.IntToInt4(source.PsychologicalTestID)
	postgresPsychologicalPerformance.StartedAt = postgres.TimeToTimestamptz(source.StartedAt)
	return postgresPsychologicalPerformance
}

func (c *ConverterImpl) PsychologicalTestDBToModel(source sqlc.PsychologicalTest) models.PsychologicalTest {
	var modelsPsychologicalTest models.PsychologicalTest
	modelsPsychologicalTest.ID = source.ID
	modelsPsychologicalTest.FirstQuestionID = postgres.Int4ToInt(source.FirstQuestionID)
	modelsPsychologicalTest.TypeID = postgres.Int4ToInt(source.TypeID)
	modelsPsychologicalTest.OwnerID = source.OwnerID
	modelsPsychologicalTest.Title = source.Title
	return modelsPsychologicalTest
}

func (c *ConverterImpl) PsychologicalTestModelToDB(source models.PsychologicalTest) sqlc.PsychologicalTest {
	var postgresPsychologicalTest sqlc.PsychologicalTest
	postgresPsychologicalTest.ID = source.ID
	postgresPsychologicalTest.FirstQuestionID = postgres.IntToInt4(source.FirstQuestionID)
	postgresPsychologicalTest.TypeID = postgres.IntToInt4(source.TypeID)
	postgresPsychologicalTest.OwnerID = source.OwnerID
	postgresPsychologicalTest.Title = source.Title
	return postgresPsychologicalTest
}

func (c *ConverterImpl) PsychologicalTestsDBToModel(source []sqlc.PsychologicalTest) []models.PsychologicalTest {
	var modelsPsychologicalTestList []models.PsychologicalTest
	if source != nil {
		modelsPsychologicalTestList = make([]models.PsychologicalTest, len(source))
		for i := 0; i < len(source); i++ {
			modelsPsychologicalTestList[i] = c.PsychologicalTestDBToModel(source[i])
		}
	}
	return modelsPsychologicalTestList
}

func (c *ConverterImpl) PsychologicalTestsModelToDB(source []models.PsychologicalTest) []sqlc.PsychologicalTest {
	var postgresPsychologicalTestList []sqlc.PsychologicalTest
	if source != nil {
		postgresPsychologicalTestList = make([]sqlc.PsychologicalTest, len(source))
		for i := 0; i < len(source); i++ {
			postgresPsychologicalTestList[i] = c.PsychologicalTestModelToDB(source[i])
		}
	}
	return postgresPsychologicalTestList
}
