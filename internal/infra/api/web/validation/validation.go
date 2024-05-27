package validation

import (
	"encoding/json"
	"errors"
	"fullcycle-auction_go/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		enTransl := ut.New(en, en)
		transl, _ = enTransl.GetTranslator("en")
		validator_en.RegisterDefaultTranslations(value, transl)
	}
}

func ValidateErr(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewNotFoundError("Invalid type error")
	} else if errors.As(validation_err, &jsonValidation) {
		errorCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			errorCauses = append(errorCauses, rest_err.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			})
		}

		return rest_err.NewBadRequestError("Invalid field values", errorCauses...)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
