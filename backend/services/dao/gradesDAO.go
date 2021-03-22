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
	if len(g.store)-1 < id {
		return nil
	}
	if g.store[id].Student == "" {
		return nil
	}
	return &g.store[id]
}

func (g *GradesStore) GetGradesByStudent(student string) []model.Grade {
	if g.store == nil {
		return make([]model.Grade, 0)
	}
	grades := make([]model.Grade, 0)

	for i := 0; i < len(g.store); i++ {
		studentFromStore := g.store[i].Student
		if studentFromStore == student {
			grades = append(grades, g.store[i])
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

	if len(g.store)-1 < id {
		return nil
	}

	if g.store[id].Student == "" {
		return nil
	}

	grade := g.store[id]
	g.store[id] = model.Grade{ID: idS}
	return &grade
}

func (g *GradesStore) EditGrade(idS string, grade model.Grade) *model.Grade {
	id, _ := strconv.Atoi(idS)

	if len(g.store)-1 < id {
		return nil
	}

	if g.store[id].Student == "" {
		return nil
	}

	g.store[id] = grade
	return &grade
}
