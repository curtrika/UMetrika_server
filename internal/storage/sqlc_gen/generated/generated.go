// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import (
	models "github.com/curtrika/UMetrika_server/internal/domain/models"
	sqlcgen "github.com/curtrika/UMetrika_server/internal/storage/sqlc_gen"
)

type ConverterImpl struct{}

func (c *ConverterImpl) PsychologicalPerfomanceDBToModel(source *sqlcgen.PsychologicalPerformance) *models.PsychologicalPerformance {
	var pModelsPsychologicalPerformance *models.PsychologicalPerformance
	if source != nil {
		var modelsPsychologicalPerformance models.PsychologicalPerformance
		modelsPsychologicalPerformance.ID = (*source).ID
		modelsPsychologicalPerformance.OwnerID = (*source).OwnerID
		modelsPsychologicalPerformance.PsychologicalTestID = sqlcgen.Int4ToInt((*source).PsychologicalTestID)
		modelsPsychologicalPerformance.StartedAt = sqlcgen.TimestamptzToTime((*source).StartedAt)
		pModelsPsychologicalPerformance = &modelsPsychologicalPerformance
	}
	return pModelsPsychologicalPerformance
}
func (c *ConverterImpl) PsychologicalPerfomanceModelToDB(source *models.PsychologicalPerformance) *sqlcgen.PsychologicalPerformance {
	var pStoragePsychologicalPerformance *sqlcgen.PsychologicalPerformance
	if source != nil {
		var storagePsychologicalPerformance sqlcgen.PsychologicalPerformance
		storagePsychologicalPerformance.ID = (*source).ID
		storagePsychologicalPerformance.OwnerID = (*source).OwnerID
		storagePsychologicalPerformance.PsychologicalTestID = sqlcgen.IntToInt4((*source).PsychologicalTestID)
		storagePsychologicalPerformance.StartedAt = sqlcgen.TimeToTimestamptz((*source).StartedAt)
		pStoragePsychologicalPerformance = &storagePsychologicalPerformance
	}
	return pStoragePsychologicalPerformance
}
func (c *ConverterImpl) PsychologicalTestDBToModel(source *sqlcgen.PsychologicalTest) *models.PsychologicalTest {
	var pModelsPsychologicalTest *models.PsychologicalTest
	if source != nil {
		var modelsPsychologicalTest models.PsychologicalTest
		modelsPsychologicalTest.ID = (*source).ID
		modelsPsychologicalTest.FirstQuestionID = sqlcgen.Int4ToInt((*source).FirstQuestionID)
		modelsPsychologicalTest.TypeID = sqlcgen.Int4ToInt((*source).TypeID)
		modelsPsychologicalTest.OwnerID = (*source).OwnerID
		modelsPsychologicalTest.Title = (*source).Title
		pModelsPsychologicalTest = &modelsPsychologicalTest
	}
	return pModelsPsychologicalTest
}
func (c *ConverterImpl) PsychologicalTestModelToDB(source *models.PsychologicalTest) *sqlcgen.PsychologicalTest {
	var pStoragePsychologicalTest *sqlcgen.PsychologicalTest
	if source != nil {
		var storagePsychologicalTest sqlcgen.PsychologicalTest
		storagePsychologicalTest.ID = (*source).ID
		storagePsychologicalTest.FirstQuestionID = sqlcgen.IntToInt4((*source).FirstQuestionID)
		storagePsychologicalTest.TypeID = sqlcgen.IntToInt4((*source).TypeID)
		storagePsychologicalTest.OwnerID = (*source).OwnerID
		storagePsychologicalTest.Title = (*source).Title
		pStoragePsychologicalTest = &storagePsychologicalTest
	}
	return pStoragePsychologicalTest
}
func (c *ConverterImpl) PsychologicalTestsDBToModel(source []sqlcgen.PsychologicalTest) []models.PsychologicalTest {
	var modelsPsychologicalTestList []models.PsychologicalTest
	if source != nil {
		modelsPsychologicalTestList = make([]models.PsychologicalTest, len(source))
		for i := 0; i < len(source); i++ {
			modelsPsychologicalTestList[i] = c.storagePsychologicalTestToModelsPsychologicalTest(source[i])
		}
	}
	return modelsPsychologicalTestList
}
func (c *ConverterImpl) PsychologicalTestsModelToDB(source []models.PsychologicalTest) []sqlcgen.PsychologicalTest {
	var storagePsychologicalTestList []sqlcgen.PsychologicalTest
	if source != nil {
		storagePsychologicalTestList = make([]sqlcgen.PsychologicalTest, len(source))
		for i := 0; i < len(source); i++ {
			storagePsychologicalTestList[i] = c.modelsPsychologicalTestToStoragePsychologicalTest(source[i])
		}
	}
	return storagePsychologicalTestList
}
func (c *ConverterImpl) modelsPsychologicalTestToStoragePsychologicalTest(source models.PsychologicalTest) sqlcgen.PsychologicalTest {
	var storagePsychologicalTest sqlcgen.PsychologicalTest
	storagePsychologicalTest.ID = source.ID
	storagePsychologicalTest.FirstQuestionID = sqlcgen.IntToInt4(source.FirstQuestionID)
	storagePsychologicalTest.TypeID = sqlcgen.IntToInt4(source.TypeID)
	storagePsychologicalTest.OwnerID = source.OwnerID
	storagePsychologicalTest.Title = source.Title
	return storagePsychologicalTest
}
func (c *ConverterImpl) storagePsychologicalTestToModelsPsychologicalTest(source sqlcgen.PsychologicalTest) models.PsychologicalTest {
	var modelsPsychologicalTest models.PsychologicalTest
	modelsPsychologicalTest.ID = source.ID
	modelsPsychologicalTest.FirstQuestionID = sqlcgen.Int4ToInt(source.FirstQuestionID)
	modelsPsychologicalTest.TypeID = sqlcgen.Int4ToInt(source.TypeID)
	modelsPsychologicalTest.OwnerID = source.OwnerID
	modelsPsychologicalTest.Title = source.Title
	return modelsPsychologicalTest
}
