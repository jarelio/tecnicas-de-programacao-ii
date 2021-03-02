package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/services"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/utils"
)

type GradesController struct {
	store services.GradesStore
}

func (c *GradesController) CleanStore() {
	var storeNil services.GradesStore
	c.store = storeNil
}

func sendHTTPSuccessResponseMessage(w http.ResponseWriter, message, data string, statusCode int) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, utils.ResultMessageAndDataToJSON(message, data))
}

func sendHTTPErrorResponseMessage(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, utils.ErrorMessageToJSON(errorMessage))
}

func (c *GradesController) GetGrades(w http.ResponseWriter, r *http.Request) {
	grades := c.store.GetGrades()

	gradesJSON, _ := json.Marshal(grades)
	sendHTTPSuccessResponseMessage(w, utils.AllGrades, string(gradesJSON), http.StatusOK)
}

func (c *GradesController) GetGrade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	grade := c.store.GetGrade(params["id"])

	if grade == nil {
		sendHTTPErrorResponseMessage(w, utils.GradeNotFound, http.StatusBadRequest)
		return
	} else {
		gradeJSON, _ := json.Marshal(grade)
		sendHTTPSuccessResponseMessage(w, utils.GradeRetrieved, string(gradeJSON), http.StatusOK)
		return
	}
}

func (c *GradesController) CreateGrade(w http.ResponseWriter, r *http.Request) {
	var grade services.Grade
	err := json.NewDecoder(r.Body).Decode(&grade)

	if err != nil {
		sendHTTPErrorResponseMessage(w, utils.InvalidParameters, http.StatusBadRequest)
		return
	}

	if grade.Student == "" || grade.Subject == "" || grade.Type == "" {
		sendHTTPErrorResponseMessage(w, utils.MissingParameters, http.StatusBadRequest)
		return
	}

	if gradeValue, _ := strconv.Atoi(grade.Value); gradeValue <= 0 {
		sendHTTPErrorResponseMessage(w, utils.ValueShouldBeGreater, http.StatusBadRequest)
		return
	}

	gradeInserted := c.store.PostGrade(grade)
	gradeJSON, _ := json.Marshal(gradeInserted)
	sendHTTPSuccessResponseMessage(w, fmt.Sprintf(utils.GradeInserted, gradeInserted.ID), string(gradeJSON), http.StatusOK)
}
func (c *GradesController) DeleteGrade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	deletedGrade := c.store.DeleteGrade(id)

	if deletedGrade == nil {
		sendHTTPErrorResponseMessage(w, utils.DeleteFailed, http.StatusBadRequest)
		return
	}

	deletedGradeJSON, _ := json.Marshal(deletedGrade)
	sendHTTPSuccessResponseMessage(w, fmt.Sprintf(utils.GradeDeleted, deletedGrade.ID), string(deletedGradeJSON), http.StatusOK)

}

func (c *GradesController) EditGrade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var grade services.Grade
	err := json.NewDecoder(r.Body).Decode(&grade)

	if err != nil {
		sendHTTPErrorResponseMessage(w, utils.InvalidParameters, http.StatusBadRequest)
		return
	}

	if grade.Student == "" || grade.Subject == "" || grade.Type == "" {
		sendHTTPErrorResponseMessage(w, utils.MissingParameters, http.StatusBadRequest)
		return
	}

	if gradeValue, _ := strconv.Atoi(grade.Value); gradeValue <= 0 {
		sendHTTPErrorResponseMessage(w, utils.ValueShouldBeGreater, http.StatusBadRequest)
		return
	}

	editedGrade := c.store.EditGrade(params["id"], grade)

	if editedGrade == nil {
		sendHTTPErrorResponseMessage(w, utils.EditFailed, http.StatusBadRequest)
		return
	}

	editedGradeJSON, _ := json.Marshal(editedGrade)
	sendHTTPSuccessResponseMessage(w, fmt.Sprintf(utils.GradeEdited, string(editedGrade.ID)), string(editedGradeJSON), http.StatusOK)

}

func (c *GradesController) GetGradesByStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	student := params["student"]
	grades := c.store.GetGradesByStudent(student)

	gradesJSON, _ := json.Marshal(grades)
	sendHTTPSuccessResponseMessage(w, utils.GradesByStudent, string(gradesJSON), http.StatusOK)
}
