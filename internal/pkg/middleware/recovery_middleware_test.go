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

	type TestCase struct {
		name            string
		responseCode    int
		responseMessage string
		funcHandler     func(c *gin.Context)
	}

	cases := []TestCase{
		{
			name:            "with error panic",
			responseCode:    http.StatusInternalServerError,
			responseMessage: "internal server error",
			funcHandler: func(c *gin.Context) {
				panic("pannic recovery")
			},
		},
		{
			name:            "without error panic",
			responseCode:    http.StatusOK,
			responseMessage: "success response",
			funcHandler: func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "success response"})
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			// SETUP
			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.Use(RecoveryMiddleware())
			router.GET("/recovery", tc.funcHandler)

			// PERFORM REQUEST
			response := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/recovery", nil)
			router.ServeHTTP(response, request)

			// PARSE RESPONSE
			var bodyJson map[string]interface{}
			respBody, _ := io.ReadAll(response.Body)
			_ = json.Unmarshal(respBody, &bodyJson)

			// CHECK ASSERT
			assert.Equal(t, tc.responseCode, response.Code)
			assert.Equal(t, tc.responseMessage, bodyJson["message"])

		})
	}
}
