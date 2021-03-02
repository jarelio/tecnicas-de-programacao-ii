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

		expectedResponse := utils.Response{
			Body:       utils.ResultMessageAndDataToJSON(utils.AllGrades, "[]"),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)
	})

	t.Run("returns all grades after some insertions", func(t *testing.T) {
		defer cleanStore()

		grades := generateGrades(2, 2)
		createGrades(router, grades)

		request := utils.NewGetGradesRequest()
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedGradesResponseJSON, _ := json.Marshal(grades)

		expectedResponse := utils.Response{
			Body:       utils.ResultMessageAndDataToJSON(utils.AllGrades, string(expectedGradesResponseJSON[:])),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)
	})
}

func TestGETGrade(t *testing.T) {

	t.Run("returns a not found grade error", func(t *testing.T) {
		defer cleanStore()
		request := utils.NewGetGradeRequest("0")
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := utils.Response{
			Body:       utils.ErrorMessageToJSON(utils.GradeNotFound),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)
	})

	t.Run("returns a specific grade", func(t *testing.T) {
		defer cleanStore()
		request := utils.NewGetGradeRequest("0")
		response := httptest.NewRecorder()

		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: "10", Student: "student_test"}
		router.ServeHTTP(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))

		router.ServeHTTP(response, request)

		gradeJSON, _ := json.Marshal(grade)

		expectedResponse := utils.Response{
			Body:       utils.ResultMessageAndDataToJSON(utils.GradeRetrieved, string(gradeJSON)),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)
	})

}

func TestPOSTGrade(t *testing.T) {

	t.Run("insert a grade with invalid parameters", func(t *testing.T) {
		defer cleanStore()
		grade := services.InvalidGrade{ID: "0", Subject: map[string]string{}, Type: "type_test", Value: "10", Student: "student_test"}
		request := utils.NewPostInvalidGradeRequest(grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := utils.Response{
			Body:       utils.ErrorMessageToJSON(utils.InvalidParameters),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("insert a grade with missing parameters", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Type: "type_test", Value: "10", Student: "student_test"}
		request := utils.NewPostGradeRequest(grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := utils.Response{
			Body:       utils.ErrorMessageToJSON(utils.MissingParameters),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("insert a grade with value less than 0", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: "-10", Student: "student_test"}
		request := utils.NewPostGradeRequest(grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := utils.Response{
			Body:       utils.ErrorMessageToJSON(utils.ValueShouldBeGreater),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("insert a grade", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: "10", Student: "student_test"}
		request := utils.NewPostGradeRequest(grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		gradeJSON, _ := json.Marshal(grade)

		expectedResponse := utils.Response{
			Body:       utils.ResultMessageAndDataToJSON(fmt.Sprintf(utils.GradeInserted, string(grade.ID)), string(gradeJSON[:])),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)

	})
}

func TestPUTGrade(t *testing.T) {

	t.Run("edit a grade with invalid parameters", func(t *testing.T) {
		defer cleanStore()
		grade := services.InvalidGrade{ID: "0", Subject: map[string]string{}, Type: "type_test", Value: "10", Student: "student_test"}
		request := utils.NewEditInvalidGradeRequest("0", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := utils.Response{
			Body:       utils.ErrorMessageToJSON(utils.InvalidParameters),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("edit a grade with missing parameters", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Type: "type_test", Value: "10", Student: "student_test"}
		request := utils.NewEditGradeRequest("0", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := utils.Response{
			Body:       utils.ErrorMessageToJSON(utils.MissingParameters),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("edit a grade with value less than 0", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: "-10", Student: "student_test"}
		request := utils.NewEditGradeRequest("0", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := utils.Response{
			Body:       utils.ErrorMessageToJSON(utils.ValueShouldBeGreater),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("edit a non existent grade", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "5", Subject: "subject_test", Type: "type_test", Value: "10", Student: "student_test"}
		request := utils.NewEditGradeRequest("5", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := utils.Response{
			Body:       utils.ErrorMessageToJSON(utils.EditFailed),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("edit a grade", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: "10", Student: "student_test"}

		request := utils.NewEditGradeRequest("0", grade)
		response := httptest.NewRecorder()

		router.ServeHTTP(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))

		router.ServeHTTP(response, request)

		gradeJSON, _ := json.Marshal(grade)

		expectedResponse := utils.Response{
			Body:       utils.ResultMessageAndDataToJSON(fmt.Sprintf(utils.GradeEdited, string(grade.ID)), string(gradeJSON[:])),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)

	})
}

func TestDELETEGrade(t *testing.T) {

	t.Run("failed to delete a grade", func(t *testing.T) {
		defer cleanStore()
		request := utils.NewDeleteGradeRequest("0")
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedResponse := utils.Response{
			Body:       utils.ErrorMessageToJSON(utils.DeleteFailed),
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)

	})

	t.Run("delete a grade successfully", func(t *testing.T) {
		defer cleanStore()
		grade := services.Grade{ID: "0", Subject: "subject_test", Type: "type_test", Value: "10", Student: "student_test"}

		request := utils.NewDeleteGradeRequest("0")
		response := httptest.NewRecorder()

		router.ServeHTTP(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))

		router.ServeHTTP(response, request)

		gradeJSON, _ := json.Marshal(grade)

		expectedResponse := utils.Response{
			Body:       utils.ResultMessageAndDataToJSON(fmt.Sprintf(utils.GradeDeleted, string(grade.ID)), string(gradeJSON[:])),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)

	})
}

func TestGETGradesByStudent(t *testing.T) {

	t.Run("get all grades by a student", func(t *testing.T) {
		defer cleanStore()

		grades := generateGrades(2, 4)
		createGrades(router, grades)

		request := utils.NewGetGradesByStudentRequest("student1")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		expectedGrades := getGradesByStudent("student1", grades)
		expectedGradesJSON, _ := json.Marshal(expectedGrades)

		expectedResponse := utils.Response{
			Body:       utils.ResultMessageAndDataToJSON(utils.GradesByStudent, string(expectedGradesJSON[:])),
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}

		utils.AssertResponse(t, response, expectedResponse)
	})
}

func generateGrades(numberOfStudents, numberOfGrades int) []services.Grade {
	grades := make([]services.Grade, 0)
	students := make([]string, 0)

	for studentNumber := 0; studentNumber < numberOfStudents; studentNumber++ {
		students = append(students, fmt.Sprintf("student%d", studentNumber))
	}

	for gradeNumber := 0; gradeNumber < numberOfGrades; gradeNumber++ {
		subject := fmt.Sprintf("subject%d", gradeNumber)
		gradeType := fmt.Sprintf("type%d", gradeNumber)
		value := "10"
		studentID := gradeNumber % numberOfStudents
		grades = append(grades, services.Grade{ID: fmt.Sprint(gradeNumber), Subject: subject, Type: gradeType, Value: value, Student: students[studentID]})
	}

	return grades
}

func createGrades(router *mux.Router, grades []services.Grade) {
	for _, grade := range grades {
		router.ServeHTTP(httptest.NewRecorder(), utils.NewPostGradeRequest(grade))
	}
}

func getGradesByStudent(student string, grades []services.Grade) []services.Grade {
	gradesByStudent := make([]services.Grade, 0)

	for _, grade := range grades {
		if grade.Student == student {
			gradesByStudent = append(gradesByStudent, grade)
		}
	}

	return gradesByStudent
}
