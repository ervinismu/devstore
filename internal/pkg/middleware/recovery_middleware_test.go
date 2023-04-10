package middleware

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRecoveryMiddleware(t *testing.T) {
	// TEST CASES
	type TestCase struct {
		Name            string
		ResponseCode    int
		ResponseMessage string
		FuncHandler     func(c *gin.Context)
	}

	cases := []TestCase{
		{
			Name:            "with error panic",
			ResponseCode:    http.StatusInternalServerError,
			ResponseMessage: "internal server error",
			FuncHandler: func(c *gin.Context) {
				panic("pannic error")
			},
		},
		{
			Name:            "without error panic",
			ResponseCode:    http.StatusOK,
			ResponseMessage: "success response",
			FuncHandler: func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "success response"})
			},
		},
	}

	// EXECUTION
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			// SETUP
			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.Use(RecoveryMiddleware())
			router.GET("/recovery", tc.FuncHandler)

			// PERFORM REQUEST
			response := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/recovery", nil)
			router.ServeHTTP(response, request)

			// PARSE RESPONSE
			var bodyJson map[string]interface{}
			respBody, _ := io.ReadAll(response.Body)
			_ = json.Unmarshal(respBody, &bodyJson)

			// CHECK ASSERT
			assert.Equal(t, tc.ResponseCode, response.Code)
			assert.Equal(t, tc.ResponseMessage, bodyJson["message"])
		})
	}
}
