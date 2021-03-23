package dao

import (
	"strconv"

	"github.com/jarelio/tecnicas-de-programacao-ii/backend/services/model"
)

type GradesStore struct {
	store []model.Grade
}

func (g *GradesStore) GetGrades() []model.Grade {
	if g.store == nil {
		return make([]model.Grade, 0)
	}
	return g.store
}

func (g *GradesStore) GetGrade(idS string) *model.Grade {
	id, _ := strconv.Atoi(idS)

	if isInvalidGradeID(g.store, id) {
		return nil
	}

	return &g.store[id]
}

func (g *GradesStore) GetGradesByStudent(student string) []model.Grade {
	grades := make([]model.Grade, 0)

	for _, grade := range g.store {
		if grade.Student == student {
			grades = append(grades, grade)
		}
	}

	return grades
}

func (g *GradesStore) PostGrade(grade model.Grade) *model.Grade {
	grade.ID = strconv.Itoa(len(g.store))
	g.store = append(g.store, grade)
	return &grade
}

func (g *GradesStore) DeleteGrade(idS string) *model.Grade {
	id, _ := strconv.Atoi(idS)

	if isInvalidGradeID(g.store, id) {
		return nil
	}

	grade := g.store[id]
	g.store[id] = model.Grade{ID: idS}
	return &grade
}

func (g *GradesStore) EditGrade(idS string, grade model.Grade) *model.Grade {
	id, _ := strconv.Atoi(idS)

	if isInvalidGradeID(g.store, id) {
		return nil
	}

	g.store[id] = grade
	return &grade
}

func isInvalidGradeID(store []model.Grade, id int) bool {
	return len(store)-1 < id || store[id].Student == ""
}
