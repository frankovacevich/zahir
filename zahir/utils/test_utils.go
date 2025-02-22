package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func AssertErrorIsNil(err error, t *testing.T) {
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func AssertError(err error, t *testing.T) {
	if err == nil {
		t.Fatalf("Error expected")
	}
}

func AssertIntEqual(expected, actual int, t *testing.T) {
	if expected != actual {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func AssertFloatEqual(expected, actual float64, t *testing.T) {
	if expected != actual {
		t.Fatalf("Expected %f, got %f", expected, actual)
	}
}

func AssertStringEqual(expected, actual string, t *testing.T) {
	if expected != actual {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func AssertPointersEqual(expected, actual interface{}, t *testing.T) {
	if expected != actual {
		t.Fatalf("Expected %p, got %p", expected, actual)
	}
}

func MakeGetRequest(url string, router *mux.Router, t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	return recorder
}

func MakePostRequest(url string, router *mux.Router, t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	return recorder
}

func GetResponseBodyAsJsonList(r *httptest.ResponseRecorder, t *testing.T) []interface{} {
	var body []interface{}
	err := json.Unmarshal(r.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("%v", err)
	}
	return body
}

func GetResponseBodyAsJsonMap(r *httptest.ResponseRecorder, t *testing.T) map[string]interface{} {
	var body map[string]interface{}
	err := json.Unmarshal(r.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("%v", err)
	}
	return body
}
