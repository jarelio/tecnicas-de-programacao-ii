package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jarelio/tecnicas-de-programacao-ii/backend/services"
)

func NewGetGradesRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/grades", nil)
	return req
}

func NewGetGradeRequest(id string) *http.Request {
	url := fmt.Sprintf("/grades/%s", id)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	return req
}

func NewPostInvalidGradeRequest(grade services.InvalidGrade) *http.Request {
	gradeJSON, _ := json.Marshal(grade)
	req, _ := http.NewRequest(http.MethodPost, "/grades", bytes.NewBuffer(gradeJSON))
	return req
}

func NewPostGradeRequest(grade services.Grade) *http.Request {
	gradeJSON, _ := json.Marshal(grade)
	req, _ := http.NewRequest(http.MethodPost, "/grades", bytes.NewBuffer(gradeJSON))
	return req
}

func NewEditInvalidGradeRequest(id string, grade services.InvalidGrade) *http.Request {
	gradeJSON, _ := json.Marshal(grade)
	url := fmt.Sprintf("/grades/%s", id)
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(gradeJSON))
	return req
}

func NewEditGradeRequest(id string, grade services.Grade) *http.Request {
	gradeJSON, _ := json.Marshal(grade)
	url := fmt.Sprintf("/grades/%s", id)
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(gradeJSON))
	return req
}

func NewDeleteGradeRequest(id string) *http.Request {
	url := fmt.Sprintf("/grades/%s", id)
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	return req
}
