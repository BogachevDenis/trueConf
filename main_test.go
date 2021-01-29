package main

import (

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)


func TestCreateUser(t *testing.T) {
	type testpair struct {
		value 		string
		wantCode 	int
	}
	var tests = []testpair{
		{`{"name":"Denis"}`,	200},
		{`{"name":"ivan"}`,		200},
		{`{"name":van"}`,		400},
	}
	for _, pair := range tests {
		e := echo.New()
		userJSON := pair.value
		req, _ := http.NewRequest(http.MethodPost, "/user", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		createUser(c)

		if rec.Code != pair.wantCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
			rec.Code, pair.wantCode)
		}
	}
}

func TestGetUserList(t *testing.T) {
	e := echo.New()
	
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	getUserList(c)
	if rec.Code != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
		rec.Code, 200)
	}
}

func TestGetUser(t *testing.T) {
	type testpair struct {
		value 		string
		wantCode 	int
	}
	var tests = []testpair{
		{"1",		200},
		{"2",		200},
		{"12452",	400},
	}

	for _, pair := range tests {
		e := echo.New()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/user/:id")
		c.SetParamNames("id")
		c.SetParamValues(pair.value)
		getUser(c)

		if rec.Code != pair.wantCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
			rec.Code, pair.wantCode)
		}
	}
}


func TestUpdateUser(t *testing.T) {
	type testpair struct {
		value 		string
		nameValue	string
		wantCode 	int
	}
	var tests = []testpair{
		{"1", 		`{"name":"kok"}`,	200},
		{"1", 		`{"name":kok"}`,	400},
		{"12313", 	`{"name":"kok"}`,	400},
	}

	for _, pair := range tests {
		e := echo.New()
		req, _ := http.NewRequest(http.MethodPut, "/", strings.NewReader(pair.nameValue))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/user/:id")
		c.SetParamNames("id")
		c.SetParamValues(pair.value)
		updateUser(c)

		if rec.Code != pair.wantCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
			rec.Code, pair.wantCode)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	type testpair struct {
		value 		string
		wantCode 	int
	}
	var tests = []testpair{
		{"1", 			200},
		{"11231231",	400},
	}

	for _, pair := range tests {
		e := echo.New()
		req, _ := http.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/user/:id")
		c.SetParamNames("id")
		c.SetParamValues(pair.value)
		deleteUser(c)

		if rec.Code != pair.wantCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
			rec.Code, pair.wantCode)
		}
	}
}