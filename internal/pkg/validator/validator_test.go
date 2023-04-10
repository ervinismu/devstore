package validator

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_Check(t *testing.T) {

	type DummyReq struct {
		Name        string `validate:"required"`
		Description string `validate:"required"`
	}

	type TestCase struct {
		Name         string
		DataNotValid bool
		ReqBody      string
	}

	cases := []TestCase{
		{
			Name:         "when name not presence",
			DataNotValid: true,
			ReqBody:      `{"description": "huhu"}`,
		},
		{
			Name:         "when description not presence",
			DataNotValid: true,
			ReqBody:      `{"name": "hihi"}`,
		},
		{
			Name:         "when name and description presence",
			DataNotValid: false,
			ReqBody:      `{"name": "hihi", "description": "huhu"}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			var req DummyReq
			_ = json.Unmarshal([]byte(tc.ReqBody), &req)
			isError := Check(&req)
			assert.Equal(t, tc.DataNotValid, isError)
		})
	}
}
