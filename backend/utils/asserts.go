package utils

import (
	"net/http/httptest"
	"testing"
)

func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Response body is wrong, got '%s' want '%s'", got, want)
	}
}

// func AssertGrades(t *testing.T, got, want []services.Grade) {
// 	t.Helper()
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("Response body is wrong, got %v want %v", got, want)
// 	}
// }

func AssertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response does not have content-type of %s, got %v", want, response.Result().Header)
	}
}

func AssertStatus(t *testing.T, code, want int) {
	t.Helper()
	if code != want {
		t.Errorf("response does not have status code of %d, got %d", want, code)
	}
}
