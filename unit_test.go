package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"Piyaporn_agnos_backend/database"
	"Piyaporn_agnos_backend/handler"
	"Piyaporn_agnos_backend/model"
)

func ModelTestCalculateNumStep(t *testing.T) {
	tests := []struct {
		passwordInput string
		expectedCount int
	}{
		{"abcABC123", 0},
		{".lfx!kk", 1},
		{"123", 2},
		{"aaa", 3},
	}

	for _, test := range tests {
		t.Run(test.passwordInput, func(t *testing.T) {
			count := model.CalculateNumStep(test.passwordInput)
			assert.Equal(t, test.expectedCount, count, "unexpected count")
		})
	}
}

func ModelTestAddData(t *testing.T) {
	db, err := database.InitDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	tests := []struct {
		passwordInput string
	}{
		{"addData001"},
		{"testData100"},
	}

	for _, test := range tests {
		t.Run(test.passwordInput, func(t *testing.T) {
			err := model.AddData(db, test.passwordInput)
			assert.NoError(t, err, "unexpected error")
		})
	}
}

func ModelTestIsValidInput(t *testing.T) {
	tests := []struct {
		input         string
		expectedValid bool
	}{
		{"abc123", true},
		{"abc!@#", false},
		{"1", true},
		{"*pwd*", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			valid := model.IsValidInput(test.input)
			assert.Equal(t, test.expectedValid, valid, "unexpected validity")
		})
	}
}

func HandlerTestAddDataHandler(t *testing.T) {
	db, err := database.InitDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	r.POST("/addData", handler.AddDataHandler(db))

	tests := []struct {
		test_pwd_input    string
		expectedStatus int
	}{
		{
			test_pwd_input:    `{"password_input": "validpassword"}`,
			expectedStatus: http.StatusOK,
		},
		{
			test_pwd_input:    `{"password_input": "invalid-password"}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.test_pwd_input, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/addData", bytes.NewBufferString(test.test_pwd_input))
			if err != nil {
				t.Fatal(err)
			}
	
			recorder := httptest.NewRecorder()
			r.ServeHTTP(recorder, req)
			assert.Equal(t, test.expectedStatus, recorder.Code, "unexpected status code")
		})
	}
}
