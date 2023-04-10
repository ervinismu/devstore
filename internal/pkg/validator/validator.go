package validator

import (
	"errors"
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
	if err != nil {
		var valErrors validator.ValidationErrors
		errors.As(err, &valErrors)
		for _, fieldError := range valErrors {
			msg := fmt.Sprintf("%s, %s", fieldError.Field(), fieldError.Error())
			log.Error(msg)
		}
		log.Error(err)
		return true
	}

	return false
}
