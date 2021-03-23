package services

import (
	"fmt"
	"strconv"

	"github.com/jarelio/tecnicas-de-programacao-ii/backend/services/dao"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/services/model"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/utils"
)

type GradesService struct {
	store dao.GradesStore
}

func (gs *GradesService) CleanStore() {
	var emptyStore dao.GradesStore
	gs.store = emptyStore
}

func (gs *GradesService) GetGrades() []model.Grade {
	grades := gs.store.GetGrades()
	return grades
}

func (gs *GradesService) GetGrade(gradeID string) (*model.Grade, error) {
	grade := gs.store.GetGrade(gradeID)

	if grade == nil {
		return nil, fmt.Errorf("%s", utils.GradeNotFound)
	}

	return grade, nil
}

func (gs *GradesService) CreateGrade(grade model.Grade) (*model.Grade, error) {

	if gradeHasEmptyValues(grade) {
		return nil, fmt.Errorf("%s", utils.MissingParameters)
	}

	if gradeValueIsLowerThanZero(grade.Value) {
		return nil, fmt.Errorf("%s", utils.ValueShouldBeGreater)
	}

	insertedGrade := gs.store.PostGrade(grade)
	return insertedGrade, nil
}

func (gs *GradesService) DeleteGrade(gradeID string) (*model.Grade, error) {

	deletedGrade := gs.store.DeleteGrade(gradeID)

	if deletedGrade == nil {
		return nil, fmt.Errorf("%s", utils.DeleteFailed)
	}

	return deletedGrade, nil
}

func (gs *GradesService) EditGrade(gradeID string, grade model.Grade) (*model.Grade, error) {

	if gradeHasEmptyValues(grade) {
		return nil, fmt.Errorf("%s", utils.MissingParameters)
	}

	if gradeValueIsLowerThanZero(grade.Value) {
		return nil, fmt.Errorf("%s", utils.ValueShouldBeGreater)
	}

	editedGrade := gs.store.EditGrade(gradeID, grade)

	if editedGrade == nil {
		return nil, fmt.Errorf("%s", utils.EditFailed)
	}

	return editedGrade, nil
}

func (gs *GradesService) GetGradesByStudent(student string) []model.Grade {
	grades := gs.store.GetGradesByStudent(student)
	return grades
}

func gradeHasEmptyValues(grade model.Grade) bool {
	return grade.Student == "" || grade.Subject == "" || grade.Type == "" || grade.Value == ""
}

func gradeValueIsLowerThanZero(value string) bool {
	gradeValue, _ := strconv.Atoi(value)
	return gradeValue <= 0
}
