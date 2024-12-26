package handler

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/Njrctr/javacode_test_golang_junior/pkg/service"
	mock_service "github.com/Njrctr/javacode_test_golang_junior/pkg/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserIdentify(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAutorization, token string)

	testTable := []struct {
		name                string
		headerName          string
		headerValue         string
		token               string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:        "Ok",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAutorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, false, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: "1",
		},
		{
			name:                "Empty Auth Header",
			headerName:          "",
			headerValue:         "Bearer token",
			token:               "token",
			mockBehavior:        func(s *mock_service.MockAutorization, token string) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"Empty auth header"}`,
		},
		{
			name:                "Empty token",
			headerName:          "Authorization",
			headerValue:         "Bearer",
			token:               "token",
			mockBehavior:        func(s *mock_service.MockAutorization, token string) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"Invalid auth header"}`,
		},
		{
			name:                "Empty Auth type",
			headerName:          "Authorization",
			headerValue:         "NotBearer token",
			token:               "token",
			mockBehavior:        func(s *mock_service.MockAutorization, token string) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"Invalid auth header"}`,
		},
		{
			name:                "Empty Token",
			headerName:          "Authorization",
			headerValue:         "Bearer ",
			mockBehavior:        func(s *mock_service.MockAutorization, token string) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"Token is empty"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAutorization(c)
			testCase.mockBehavior(auth, testCase.token)

			services := &service.Service{Autorization: auth}
			handler := NewHandler(services)

			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.GET("/protected", handler.userIdentify, func(ctx *gin.Context) {
				id, _ := ctx.Get(userCtx)
				ctx.String(200, fmt.Sprintf("%d", id.(int)))
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}
