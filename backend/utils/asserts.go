package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	Header     http.Header
	Body       string
	StatusCode int
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Response body is wrong, got '%s' want '%s'", got, want)
	}
}

func assertContentType(t *testing.T, responseContentType, expectedContentType string) {
	t.Helper()
	if responseContentType != expectedContentType {
		t.Errorf("response does not have content-type of %s, got %s", expectedContentType, responseContentType)
	}
}

func assertStatus(t *testing.T, code, want int) {
	t.Helper()
	if code != want {
		t.Errorf("response does not have status code of %d, got %d", want, code)
	}
}

func AssertResponse(t *testing.T, response *httptest.ResponseRecorder, expectedResponse Response) {
	t.Helper()

	assertResponseBody(t, response.Body.String(), expectedResponse.Body)
	assertStatus(t, response.Code, expectedResponse.StatusCode)
	assertContentType(t, response.Result().Header.Get("content-type"), expectedResponse.Header.Get("content-type"))
}
