package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Check(value interface{}) bool {
	err := validate.Struct(value)
	log.Error(fmt.Errorf("validation error : %w", err))
	return err != nil
}
