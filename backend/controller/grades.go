package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/services"
)

type GradesController struct {
	store services.GradesStore
}

func (c *GradesController) CleanStore() {
	var storeNil services.GradesStore
	c.store = storeNil
}

func sendOKResponseMessage(w http.ResponseWriter, message, data string) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	var returnJSON string
	if data == "" {
		returnJSON = fmt.Sprintf(`{"result": {"message": "%s"}}`, message)
	} else {
		returnJSON = fmt.Sprintf(`{"result": {"message": "%s", "data": %s}}`, message, data)
	}
	fmt.Fprintln(w, returnJSON)
}

func sendBadRequestResponseMessage(w http.ResponseWriter, errorString string) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, fmt.Sprintf(`{"error": "%s"}`, errorString))
}

func (c *GradesController) GetGrades(w http.ResponseWriter, r *http.Request) {
	grades := c.store.GetGrades()

	gradesJSON, _ := json.Marshal(grades)
	sendOKResponseMessage(w, "Successfully retrieved all grades", string(gradesJSON))
}

func (c *GradesController) GetGrade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	grade := c.store.GetGrade(params["id"])

	if grade == nil {
		sendBadRequestResponseMessage(w, "Grade not found")
		return
	} else {
		gradeJSON, _ := json.Marshal(grade)
		sendOKResponseMessage(w, "Successfully retrieved the grade", string(gradeJSON))
		return
	}
}

func (c *GradesController) CreateGrade(w http.ResponseWriter, r *http.Request) {
	var grade services.Grade
	err := json.NewDecoder(r.Body).Decode(&grade)

	if err != nil {
		sendBadRequestResponseMessage(w, "Invalid parameters")
		return
	}

	if grade.Student == "" || grade.Subject == "" || grade.Type == "" {
		sendBadRequestResponseMessage(w, "Missing parameters")
		return
	}

	if gradeValue, _ := strconv.Atoi(grade.Value); gradeValue <= 0 {
		sendBadRequestResponseMessage(w, "Grade value should be greater than zero")
		return
	}

	gradeInserted := c.store.PostGrade(grade)
	gradeJSON, _ := json.Marshal(gradeInserted)
	sendOKResponseMessage(w, fmt.Sprintf("Successfully inserted grade with ID %s", gradeInserted.ID), string(gradeJSON))
}
func (c *GradesController) DeleteGrade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	deletedGrade := c.store.DeleteGrade(id)

	if deletedGrade == nil {
		sendBadRequestResponseMessage(w, "Failed to delete the grade")
		return
	}

	deletedGradeJSON, _ := json.Marshal(deletedGrade)
	sendOKResponseMessage(w, fmt.Sprintf("Successfully deleted grade with ID %s", deletedGrade.ID), string(deletedGradeJSON))

}

func (c *GradesController) EditGrade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var grade services.Grade
	err := json.NewDecoder(r.Body).Decode(&grade)

	if err != nil {
		sendBadRequestResponseMessage(w, "Invalid parameters")
		return
	}
	if grade.Student == "" || grade.Subject == "" || grade.Type == "" || grade.ID == "" {
		sendBadRequestResponseMessage(w, "Missing parameters")
		return
	}

	if gradeValue, _ := strconv.Atoi(grade.Value); gradeValue <= 0 {
		sendBadRequestResponseMessage(w, "Grade value should be greater than zero")
		return
	}

	editedGrade := c.store.EditGrade(params["id"], grade)

	if editedGrade == nil {
		sendBadRequestResponseMessage(w, "Failed to edit the grade")
		return
	}

	editedGradeJSON, _ := json.Marshal(editedGrade)
	sendOKResponseMessage(w, fmt.Sprintf("Successfully edited grade with ID %s", string(editedGrade.ID)), string(editedGradeJSON))

}

func (c *GradesController) GetGradesByStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	student := params["student"]
	grades := c.store.GetGradesByStudent(student)

	gradesJSON, _ := json.Marshal(grades)
	sendOKResponseMessage(w, "Successfully retrieved all the grades by student", string(gradesJSON))
}
