package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/services"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/services/model"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/utils"
)

const (
	//Parameter used in received requests to identify requested grade
	gradeIDParam = "id"
)

type GradesController struct {
	GradesService services.GradesService
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
	grades := c.GradesService.GetGrades()
	gradesJSON, _ := json.Marshal(grades)
	sendHTTPSuccessResponseMessage(w, utils.AllGrades, string(gradesJSON), http.StatusOK)
}

func (c *GradesController) GetGrade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	grade, err := c.GradesService.GetGrade(params[gradeIDParam])

	if err != nil {
		sendHTTPErrorResponseMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	gradeJSON, _ := json.Marshal(grade)
	sendHTTPSuccessResponseMessage(w, utils.GradeRetrieved, string(gradeJSON), http.StatusOK)
}

func (c *GradesController) CreateGrade(w http.ResponseWriter, r *http.Request) {
	var grade model.Grade
	err := json.NewDecoder(r.Body).Decode(&grade)

	if err != nil {
		sendHTTPErrorResponseMessage(w, utils.InvalidParameters, http.StatusBadRequest)
		return
	}

	gradeInserted, err := c.GradesService.CreateGrade(grade)

	if err != nil {
		sendHTTPErrorResponseMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	gradeJSON, _ := json.Marshal(gradeInserted)
	sendHTTPSuccessResponseMessage(w, fmt.Sprintf(utils.GradeInserted, gradeInserted.ID), string(gradeJSON), http.StatusOK)
}

func (c *GradesController) DeleteGrade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	deletedGrade, err := c.GradesService.DeleteGrade(params[gradeIDParam])

	if err != nil {
		sendHTTPErrorResponseMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	deletedGradeJSON, _ := json.Marshal(deletedGrade)
	sendHTTPSuccessResponseMessage(w, fmt.Sprintf(utils.GradeDeleted, deletedGrade.ID), string(deletedGradeJSON), http.StatusOK)
}

func (c *GradesController) EditGrade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var grade model.Grade
	err := json.NewDecoder(r.Body).Decode(&grade)

	if err != nil {
		sendHTTPErrorResponseMessage(w, utils.InvalidParameters, http.StatusBadRequest)
		return
	}

	editedGrade, err := c.GradesService.EditGrade(params[gradeIDParam], grade)

	if err != nil {
		sendHTTPErrorResponseMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	editedGradeJSON, _ := json.Marshal(editedGrade)
	sendHTTPSuccessResponseMessage(w, fmt.Sprintf(utils.GradeEdited, string(editedGrade.ID)), string(editedGradeJSON), http.StatusOK)
}

func (c *GradesController) GetGradesByStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	student := params["student"]
	grades := c.GradesService.GetGradesByStudent(student)

	gradesJSON, _ := json.Marshal(grades)
	sendHTTPSuccessResponseMessage(w, utils.GradesByStudent, string(gradesJSON), http.StatusOK)
}
