package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/Njrctr/javacode_test_golang_junior/pkg/service"
	mock_service "github.com/Njrctr/javacode_test_golang_junior/pkg/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandlerSignUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAutorization, user models.SignUpInput)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           models.SignUpInput
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username":"testUsername","password":"testPassword"}`,
			inputUser: models.SignUpInput{
				Username: "testUsername",
				Password: "testPassword",
			},
			mockBehavior: func(s *mock_service.MockAutorization, user models.SignUpInput) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "Empty fields",
			inputBody:           `{"username":"testUsername"}`,
			mockBehavior:        func(s *mock_service.MockAutorization, user models.SignUpInput) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:                "Invalid Input Data",
			inputBody:           `{"username":"testUsername","password":123}`,
			mockBehavior:        func(s *mock_service.MockAutorization, user models.SignUpInput) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"username":"testUsername","password":"testPassword"}`,
			inputUser: models.SignUpInput{
				Username: "testUsername",
				Password: "testPassword",
			},
			mockBehavior: func(s *mock_service.MockAutorization, user models.SignUpInput) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAutorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Autorization: auth}
			handler := NewHandler(services)

			//TEST Server
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			//TEST Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())

		})
	}
}
