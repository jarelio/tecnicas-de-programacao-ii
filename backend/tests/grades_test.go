package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarelio/tecnicas-de-programacao-ii/backend/controller"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/services"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/utils"
)

func TestGETGrades(t *testing.T) {
	controller := &controller.GradesController{}

	t.Run("returns all grades before any insertions", func(t *testing.T) {
		request := utils.NewGetGradesRequest()
		response := httptest.NewRecorder()

		controller.GetGrades(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ResultMessageAndData(utils.AllGrades, "[]"))
		utils.AssertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns all grades after some insertions", func(t *testing.T) {
		grade := services.Grade{Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"}

		request := utils.NewGetGradesRequest()
		response := httptest.NewRecorder()

		controller.CreateGrade(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))
		controller.CreateGrade(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))

		controller.GetGrades(response, request)

		want := []services.Grade{
			{ID: "0", Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"},
			{ID: "1", Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"},
		}
		wantJSON, _ := json.Marshal(want)
		utils.AssertResponseBody(t, response.Body.String(), utils.ResultMessageAndData(utils.AllGrades, string(wantJSON[:])))
		utils.AssertStatus(t, response.Code, http.StatusOK)
	})
}

func TestGETGrade(t *testing.T) {
	controller := &controller.GradesController{}

	t.Run("returns a not found grade error", func(t *testing.T) {
		request := utils.NewGetGradeRequest("0")
		response := httptest.NewRecorder()

		controller.GetGrade(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.NotFoundGrade))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("returns a specific grade", func(t *testing.T) {
		request := utils.NewGetGradeRequest("0")
		response := httptest.NewRecorder()

		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"}
		controller.CreateGrade(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))

		controller.GetGrade(response, request)

		gradeJSON, _ := json.Marshal(grade)
		utils.AssertResponseBody(t, response.Body.String(), utils.ResultMessageAndData(utils.GradeRetrieved, string(gradeJSON)))
		utils.AssertStatus(t, response.Code, http.StatusOK)
	})

}

func TestPOSTGrade(t *testing.T) {
	controller := &controller.GradesController{}

	t.Run("insert a grade with invalid parameters", func(t *testing.T) {
		grade := services.InvalidGrade{ID: "0", Subject: map[string]string{}, Type: "type_test", Value: 10, Student: "student_test"}
		request := utils.NewPostInvalidGradeRequest(grade)
		response := httptest.NewRecorder()

		controller.CreateGrade(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.InvalidParameters))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("insert a grade with missing parameters", func(t *testing.T) {
		grade := services.Grade{ID: "0", Type: "type_test", Value: 10, Student: "student_test"}
		request := utils.NewPostGradeRequest(grade)
		response := httptest.NewRecorder()

		controller.CreateGrade(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.MissingParameters))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("insert a grade with value less than 0", func(t *testing.T) {
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: -10, Student: "student_test"}
		request := utils.NewPostGradeRequest(grade)
		response := httptest.NewRecorder()

		controller.CreateGrade(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.ValueShouldBeGreater))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("insert a grade", func(t *testing.T) {
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"}
		request := utils.NewPostGradeRequest(grade)
		response := httptest.NewRecorder()

		controller.CreateGrade(response, request)

		gradeJSON, _ := json.Marshal(grade)
		utils.AssertResponseBody(t, response.Body.String(), utils.ResultMessageAndData(utils.GradeInserted, string(gradeJSON[:])))
		utils.AssertStatus(t, response.Code, http.StatusOK)
	})
}

func TestPUTGrade(t *testing.T) {
	controller := &controller.GradesController{}

	t.Run("edit a grade with invalid parameters", func(t *testing.T) {
		grade := services.InvalidGrade{ID: "0", Subject: map[string]string{}, Type: "type_test", Value: 10, Student: "student_test"}
		request := utils.NewEditInvalidGradeRequest("0", grade)
		response := httptest.NewRecorder()

		controller.EditGrade(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.InvalidParameters))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("edit a grade with missing parameters", func(t *testing.T) {
		grade := services.Grade{ID: "0", Type: "type_test", Value: 10, Student: "student_test"}
		request := utils.NewEditGradeRequest("0", grade)
		response := httptest.NewRecorder()

		controller.EditGrade(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.MissingParameters))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("edit a grade with value less than 0", func(t *testing.T) {
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: -10, Student: "student_test"}
		request := utils.NewEditGradeRequest("0", grade)
		response := httptest.NewRecorder()

		controller.EditGrade(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.ValueShouldBeGreater))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("edit a non existent grade", func(t *testing.T) {
		grade := services.Grade{ID: "5", Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"}
		request := utils.NewEditGradeRequest("5", grade)
		response := httptest.NewRecorder()

		controller.EditGrade(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.EditFailed))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("edit a grade", func(t *testing.T) {
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"}

		request := utils.NewEditGradeRequest("0", grade)
		response := httptest.NewRecorder()

		controller.CreateGrade(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))

		controller.EditGrade(response, request)

		gradeJSON, _ := json.Marshal(grade)
		utils.AssertResponseBody(t, response.Body.String(), utils.ResultMessageAndData(utils.GradeEdited, string(gradeJSON[:])))
		utils.AssertStatus(t, response.Code, http.StatusOK)
	})
}

func TestDELETEGrade(t *testing.T) {
	controller := &controller.GradesController{}

	t.Run("failed to delete a grade", func(t *testing.T) {
		request := utils.NewDeleteGradeRequest("0")
		response := httptest.NewRecorder()

		controller.DeleteGrade(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.DeleteFailed))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("delete a grade successfully", func(t *testing.T) {
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"}

		request := utils.NewDeleteGradeRequest("0")
		response := httptest.NewRecorder()

		controller.CreateGrade(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))

		controller.DeleteGrade(response, request)

		gradeJSON, _ := json.Marshal(grade)
		utils.AssertResponseBody(t, response.Body.String(), utils.ResultMessageAndData(utils.GradeDeleted, string(gradeJSON[:])))
		utils.AssertStatus(t, response.Code, http.StatusOK)
	})
}
