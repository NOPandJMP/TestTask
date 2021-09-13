package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	rw := httptest.NewRecorder()

	jsonString := "{'pcName':'test' , 'username': 'testuser' , 'networkAddr': 'localhost'}"

	req := httptest.NewRequest(http.MethodPost, "/api/v1/workspace", bytes.NewBufferString(jsonString))

	// Вызов handler IndexFunc
	WorkSpace(rw, req)

	if rw.Result().StatusCode != http.StatusOK {
		t.Errorf("Request code %v not equal code 200", rw.Code)
	}
	if rw.Body.String() != "Successfully adding user information" {
		t.Errorf(`Request body "%v" not equal body`, rw.Body.String())
	}
}
