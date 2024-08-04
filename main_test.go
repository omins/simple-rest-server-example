package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)

	assert.Equal(2, len(list))
	assert.Equal("aaa", list[0].Name)
	assert.Equal("bbb", list[1].Name)
}

func TestStudentGetHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)

	var student Student
	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	err := json.NewDecoder(res.Body).Decode(&student)

	assert.Nil(err)
	assert.Equal(1, student.Id)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusNotFound, res.Code)
}

func TestStudentPostHandler(t *testing.T) {
	var student Student
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students", strings.NewReader(`{"Id": 0, "Name" :"ccc", "Age":20, "Score":99}`))

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	err := json.NewDecoder(res.Body).Decode(&student)

	assert.Nil(err)
	assert.Equal("ccc", student.Name)
}

func TestStudentDeleteHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/students/1", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/1", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusNotFound, res.Code)
}
