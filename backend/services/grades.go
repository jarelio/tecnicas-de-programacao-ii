package services

import (
	"strconv"
)

type Grade struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	Type    string `json:"type"`
	Value   string `json:"value"`
	Student string `json:"student"`
}

type InvalidGrade struct {
	ID      string            `json:"id"`
	Subject map[string]string `json:"subject"`
	Type    string            `json:"type"`
	Value   string            `json:"value"`
	Student string            `json:"student"`
}

type GradesStore struct {
	store []Grade
}

func (g *GradesStore) GetGrades() []Grade {
	if g.store == nil {
		return make([]Grade, 0)
	}
	return g.store
}

func (g *GradesStore) GetGrade(idS string) *Grade {
	id, _ := strconv.Atoi(idS)
	if len(g.store)-1 < id {
		return nil
	}
	if g.store[id].Student == "" {
		return nil
	}
	return &g.store[id]
}

func (g *GradesStore) GetGradesByStudent(student string) []Grade {
	if g.store == nil {
		return make([]Grade, 0)
	}
	grades := make([]Grade, 0)

	for i := 0; i < len(g.store); i++ {
		studentFromStore := g.store[i].Student
		if studentFromStore == student {
			grades = append(grades, g.store[i])
		}
	}
	return grades
}

func (g *GradesStore) PostGrade(grade Grade) *Grade {
	grade.ID = strconv.Itoa(len(g.store))
	g.store = append(g.store, grade)
	return &grade
}

func (g *GradesStore) DeleteGrade(idS string) *Grade {
	id, _ := strconv.Atoi(idS)

	if len(g.store)-1 < id {
		return nil
	}

	if g.store[id].Student == "" {
		return nil
	}
	grade := g.store[id]
	g.store[id] = Grade{ID: idS}
	return &grade
}

func (g *GradesStore) EditGrade(idS string, grade Grade) *Grade {
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
