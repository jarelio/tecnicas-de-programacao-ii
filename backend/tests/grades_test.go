package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/controller"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/services"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/utils"
)

var router = mux.NewRouter()
var controllerInstance = &controller.GradesController{}

func TestMain(m *testing.M) {

	router.HandleFunc("/grades", controllerInstance.GetGrades).Methods("GET")
	router.HandleFunc("/grades", controllerInstance.CreateGrade).Methods("POST")
	router.HandleFunc("/grades/{id:[0-9]+}", controllerInstance.GetGrade).Methods("GET")
	router.HandleFunc("/grades/{id:[0-9]+}", controllerInstance.DeleteGrade).Methods("DELETE")
	router.HandleFunc("/grades/{id:[0-9]+}", controllerInstance.EditGrade).Methods("PUT")
	router.HandleFunc("/grades/student/{student:[a-zA-Z0-9_-]+}", controllerInstance.GetGradesByStudent).Methods("GET")

	code := m.Run()

	os.Exit(code)
}

func cleanStore() {
	controllerInstance.CleanStore()
}

func TestGETGrades(t *testing.T) {

	t.Run("returns all grades before any insertions", func(t *testing.T) {
		defer cleanStore()
		request := utils.NewGetGradesRequest()
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ResultMessageAndData(utils.AllGrades, "[]"))
		utils.AssertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns all grades after some insertions", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"}

		request := utils.NewGetGradesRequest()
		response := httptest.NewRecorder()

		router.ServeHTTP(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))
		router.ServeHTTP(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))

		router.ServeHTTP(response, request)

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

	t.Run("returns a not found grade error", func(t *testing.T) {
		defer cleanStore()
		request := utils.NewGetGradeRequest("0")
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.NotFoundGrade))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("returns a specific grade", func(t *testing.T) {
		defer cleanStore()
		request := utils.NewGetGradeRequest("0")
		response := httptest.NewRecorder()

		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"}
		router.ServeHTTP(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))

		router.ServeHTTP(response, request)

		gradeJSON, _ := json.Marshal(grade)
		utils.AssertResponseBody(t, response.Body.String(), utils.ResultMessageAndData(utils.GradeRetrieved, string(gradeJSON)))
		utils.AssertStatus(t, response.Code, http.StatusOK)
	})

}

func TestPOSTGrade(t *testing.T) {

	t.Run("insert a grade with invalid parameters", func(t *testing.T) {
		defer cleanStore()
		grade := services.InvalidGrade{ID: "0", Subject: map[string]string{}, Type: "type_test", Value: 10, Student: "student_test"}
		request := utils.NewPostInvalidGradeRequest(grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.InvalidParameters))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("insert a grade with missing parameters", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Type: "type_test", Value: 10, Student: "student_test"}
		request := utils.NewPostGradeRequest(grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.MissingParameters))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("insert a grade with value less than 0", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: -10, Student: "student_test"}
		request := utils.NewPostGradeRequest(grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.ValueShouldBeGreater))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("insert a grade", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"}
		request := utils.NewPostGradeRequest(grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		gradeJSON, _ := json.Marshal(grade)
		utils.AssertResponseBody(t, response.Body.String(), utils.ResultMessageAndData(utils.GradeInserted, string(gradeJSON[:])))
		utils.AssertStatus(t, response.Code, http.StatusOK)
	})
}

func TestPUTGrade(t *testing.T) {

	t.Run("edit a grade with invalid parameters", func(t *testing.T) {
		defer cleanStore()
		grade := services.InvalidGrade{ID: "0", Subject: map[string]string{}, Type: "type_test", Value: 10, Student: "student_test"}
		request := utils.NewEditInvalidGradeRequest("0", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.InvalidParameters))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("edit a grade with missing parameters", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Type: "type_test", Value: 10, Student: "student_test"}
		request := utils.NewEditGradeRequest("0", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.MissingParameters))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("edit a grade with value less than 0", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: -10, Student: "student_test"}
		request := utils.NewEditGradeRequest("0", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.ValueShouldBeGreater))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("edit a non existent grade", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "5", Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"}
		request := utils.NewEditGradeRequest("5", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.EditFailed))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("edit a grade", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"}

		request := utils.NewEditGradeRequest("0", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))

		router.ServeHTTP(response, request)

		gradeJSON, _ := json.Marshal(grade)
		utils.AssertResponseBody(t, response.Body.String(), utils.ResultMessageAndData(utils.GradeEdited, string(gradeJSON[:])))
		utils.AssertStatus(t, response.Code, http.StatusOK)
	})
}

func TestDELETEGrade(t *testing.T) {

	t.Run("failed to delete a grade", func(t *testing.T) {
		defer cleanStore()
		request := utils.NewDeleteGradeRequest("0")
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), utils.ErrorMessage(utils.DeleteFailed))
		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("delete a grade successfully", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: 10, Student: "student_test"}

		request := utils.NewDeleteGradeRequest("0")
		response := httptest.NewRecorder()

		router.ServeHTTP(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))

		router.ServeHTTP(response, request)

		gradeJSON, _ := json.Marshal(grade)

		utils.AssertResponseBody(t, response.Body.String(), utils.ResultMessageAndData(utils.GradeDeleted, string(gradeJSON[:])))
		utils.AssertStatus(t, response.Code, http.StatusOK)
	})
}

func TestGETGradesByStudent(t *testing.T) {

	t.Run("get all grades by a student", func(t *testing.T) {
		defer cleanStore()
		grade1 := services.Grade{ID: "0", Subject: "subject1", Type: "type1", Value: 10, Student: "student1"}
		grade2 := services.Grade{ID: "1", Subject: "subject2", Type: "type2", Value: 10, Student: "student2"}
		grade3 := services.Grade{ID: "2", Subject: "subject1", Type: "type1", Value: 10, Student: "student1"}
		grade4 := services.Grade{ID: "3", Subject: "subject2", Type: "type2", Value: 10, Student: "student2"}

		router.ServeHTTP(httptest.NewRecorder(), utils.NewPostGradeRequest(grade1))
		router.ServeHTTP(httptest.NewRecorder(), utils.NewPostGradeRequest(grade2))
		router.ServeHTTP(httptest.NewRecorder(), utils.NewPostGradeRequest(grade3))
		router.ServeHTTP(httptest.NewRecorder(), utils.NewPostGradeRequest(grade4))

		request := utils.NewGetGradesByStudentRequest("student1")
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		want := []services.Grade{
			{ID: "0", Subject: "subject1", Type: "type1", Value: 10, Student: "student1"},
			{ID: "2", Subject: "subject1", Type: "type1", Value: 10, Student: "student1"},
		}
		wantJSON, _ := json.Marshal(want)
		utils.AssertResponseBody(t, response.Body.String(), utils.ResultMessageAndData(utils.GradesByStudent, string(wantJSON[:])))
		utils.AssertStatus(t, response.Code, http.StatusOK)
	})
}
