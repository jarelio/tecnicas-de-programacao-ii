package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/controller"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/services/model"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/tests/testutils"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/utils"
)

var router = mux.NewRouter()
var gradesController = &controller.GradesController{}

func TestMain(m *testing.M) {

	router.HandleFunc("/grades", gradesController.GetGrades).Methods("GET")
	router.HandleFunc("/grades", gradesController.CreateGrade).Methods("POST")
	router.HandleFunc("/grades/{id:[0-9]+}", gradesController.GetGrade).Methods("GET")
	router.HandleFunc("/grades/{id:[0-9]+}", gradesController.DeleteGrade).Methods("DELETE")
	router.HandleFunc("/grades/{id:[0-9]+}", gradesController.EditGrade).Methods("PUT")
	router.HandleFunc("/grades/student/{student:[a-zA-Z0-9_-]+}", gradesController.GetGradesByStudent).Methods("GET")

	code := m.Run()

	os.Exit(code)
}

func cleanStore() {
	gradesController.GradesService.CleanStore()
}

func TestGETGrades(t *testing.T) {

	t.Run("returns all grades before any insertions", func(t *testing.T) {
		defer cleanStore()
		request := testutils.NewGetGradesRequest()
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := testutils.Response{
			Body:       utils.ResultMessageAndDataToJSON(utils.AllGrades, "[]"),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)
	})

	t.Run("returns all grades after some insertions", func(t *testing.T) {
		defer cleanStore()

		grades := generateGrades(2, 2)
		createGrades(router, grades)

		request := testutils.NewGetGradesRequest()
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedGradesResponseJSON, _ := json.Marshal(grades)

		expectedResponse := testutils.Response{
			Body:       utils.ResultMessageAndDataToJSON(utils.AllGrades, string(expectedGradesResponseJSON[:])),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)
	})
}

func TestGETGrade(t *testing.T) {

	t.Run("returns a not found grade error", func(t *testing.T) {
		defer cleanStore()
		request := testutils.NewGetGradeRequest("0")
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := testutils.Response{
			Body:       utils.ErrorMessageToJSON(utils.GradeNotFound),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)
	})

	t.Run("returns a specific grade", func(t *testing.T) {
		defer cleanStore()
		request := testutils.NewGetGradeRequest("0")
		response := httptest.NewRecorder()

		grade := model.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: "10", Student: "student_test"}
		router.ServeHTTP(httptest.NewRecorder(), testutils.NewPostGradeRequest(grade))

		router.ServeHTTP(response, request)

		gradeJSON, _ := json.Marshal(grade)

		expectedResponse := testutils.Response{
			Body:       utils.ResultMessageAndDataToJSON(utils.GradeRetrieved, string(gradeJSON)),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)
	})

}

func TestPOSTGrade(t *testing.T) {

	t.Run("insert a grade with invalid parameters", func(t *testing.T) {
		defer cleanStore()
		grade := model.InvalidGrade{ID: "0", Subject: map[string]string{}, Type: "type_test", Value: "10", Student: "student_test"}
		request := testutils.NewPostInvalidGradeRequest(grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := testutils.Response{
			Body:       utils.ErrorMessageToJSON(utils.InvalidParameters),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("insert a grade with missing parameters", func(t *testing.T) {
		defer cleanStore()
		grade := model.Grade{ID: "0", Type: "type_test", Value: "10", Student: "student_test"}
		request := testutils.NewPostGradeRequest(grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := testutils.Response{
			Body:       utils.ErrorMessageToJSON(utils.MissingParameters),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("insert a grade with value less than 0", func(t *testing.T) {
		defer cleanStore()
		grade := model.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: "-10", Student: "student_test"}
		request := testutils.NewPostGradeRequest(grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := testutils.Response{
			Body:       utils.ErrorMessageToJSON(utils.ValueShouldBeGreater),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("insert a grade", func(t *testing.T) {
		defer cleanStore()
		grade := model.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: "10", Student: "student_test"}
		request := testutils.NewPostGradeRequest(grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		gradeJSON, _ := json.Marshal(grade)

		expectedResponse := testutils.Response{
			Body:       utils.ResultMessageAndDataToJSON(fmt.Sprintf(utils.GradeInserted, string(grade.ID)), string(gradeJSON[:])),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)

	})
}

func TestPUTGrade(t *testing.T) {

	t.Run("edit a grade with invalid parameters", func(t *testing.T) {
		defer cleanStore()
		grade := model.InvalidGrade{ID: "0", Subject: map[string]string{}, Type: "type_test", Value: "10", Student: "student_test"}
		request := testutils.NewEditInvalidGradeRequest("0", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := testutils.Response{
			Body:       utils.ErrorMessageToJSON(utils.InvalidParameters),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("edit a grade with missing parameters", func(t *testing.T) {
		defer cleanStore()
		grade := model.Grade{ID: "0", Type: "type_test", Value: "10", Student: "student_test"}
		request := testutils.NewEditGradeRequest("0", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := testutils.Response{
			Body:       utils.ErrorMessageToJSON(utils.MissingParameters),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("edit a grade with value less than 0", func(t *testing.T) {
		defer cleanStore()
		grade := model.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: "-10", Student: "student_test"}
		request := testutils.NewEditGradeRequest("0", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := testutils.Response{
			Body:       utils.ErrorMessageToJSON(utils.ValueShouldBeGreater),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("edit a non existent grade", func(t *testing.T) {
		defer cleanStore()
		grade := model.Grade{ID: "5", Subject: "subject_test", Type: "type_test", Value: "10", Student: "student_test"}
		request := testutils.NewEditGradeRequest("5", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := testutils.Response{
			Body:       utils.ErrorMessageToJSON(utils.EditFailed),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("edit a grade", func(t *testing.T) {
		defer cleanStore()
		grade := model.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: "10", Student: "student_test"}

		request := testutils.NewEditGradeRequest("0", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(httptest.NewRecorder(), testutils.NewPostGradeRequest(grade))

		router.ServeHTTP(response, request)

		gradeJSON, _ := json.Marshal(grade)

		expectedResponse := testutils.Response{
			Body:       utils.ResultMessageAndDataToJSON(fmt.Sprintf(utils.GradeEdited, string(grade.ID)), string(gradeJSON[:])),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)

	})
}

func TestDELETEGrade(t *testing.T) {

	t.Run("failed to delete a grade", func(t *testing.T) {
		defer cleanStore()
		request := testutils.NewDeleteGradeRequest("0")
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := testutils.Response{
			Body:       utils.ErrorMessageToJSON(utils.DeleteFailed),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("delete a grade successfully", func(t *testing.T) {
		defer cleanStore()
		grade := model.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: "10", Student: "student_test"}

		request := testutils.NewDeleteGradeRequest("0")
		response := httptest.NewRecorder()

		router.ServeHTTP(httptest.NewRecorder(), testutils.NewPostGradeRequest(grade))

		router.ServeHTTP(response, request)

		gradeJSON, _ := json.Marshal(grade)

		expectedResponse := testutils.Response{
			Body:       utils.ResultMessageAndDataToJSON(fmt.Sprintf(utils.GradeDeleted, string(grade.ID)), string(gradeJSON[:])),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)

	})
}

func TestGETGradesByStudent(t *testing.T) {

	t.Run("get all grades by a student when student has no grades", func(t *testing.T) {
		defer cleanStore()

		request := testutils.NewGetGradesByStudentRequest("student1")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		expectedResponse := testutils.Response{
			Body:       utils.ResultMessageAndDataToJSON(utils.GradesByStudent, "[]"),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)
	})

	t.Run("get all grades by a student", func(t *testing.T) {
		defer cleanStore()

		grades := generateGrades(2, 4)
		createGrades(router, grades)

		request := testutils.NewGetGradesByStudentRequest("student1")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		expectedGrades := getGradesByStudent("student1", grades)
		expectedGradesJSON, _ := json.Marshal(expectedGrades)

		expectedResponse := testutils.Response{
			Body:       utils.ResultMessageAndDataToJSON(utils.GradesByStudent, string(expectedGradesJSON[:])),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		testutils.AssertResponse(t, response, expectedResponse)
	})
}

func generateGrades(numberOfStudents, numberOfGrades int) []model.Grade {
	grades := make([]model.Grade, 0)
	students := make([]string, 0)

	for studentNumber := 0; studentNumber < numberOfStudents; studentNumber++ {
		students = append(students, fmt.Sprintf("student%d", studentNumber))
	}

	for gradeNumber := 0; gradeNumber < numberOfGrades; gradeNumber++ {
		subject := fmt.Sprintf("subject%d", gradeNumber)
		gradeType := fmt.Sprintf("type%d", gradeNumber)
		value := "10"
		studentID := gradeNumber % numberOfStudents
		grades = append(grades, model.Grade{ID: fmt.Sprint(gradeNumber), Subject: subject, Type: gradeType, Value: value, Student: students[studentID]})
	}

	return grades
}

func createGrades(router *mux.Router, grades []model.Grade) {
	for _, grade := range grades {
		router.ServeHTTP(httptest.NewRecorder(), testutils.NewPostGradeRequest(grade))
	}
}

func getGradesByStudent(student string, grades []model.Grade) []model.Grade {
	gradesByStudent := make([]model.Grade, 0)

	for _, grade := range grades {
		if grade.Student == student {
			gradesByStudent = append(gradesByStudent, grade)
		}
	}

	return gradesByStudent
}
