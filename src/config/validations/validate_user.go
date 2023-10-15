package validations

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		transl, _ = uni.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(v, transl)
	}
}

func ValidateUserError(validationErr error) *ierrors.TError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return ierrors.NewBadRequestError()
	} else if errors.As(validationErr, &jsonValidationError) {
		errorCauses := []ierrors.TCause{}

		for _, err := range validationErr.(validator.ValidationErrors) {
			errorCauses = append(errorCauses, ierrors.TCause{
				Field:   err.Field(),
				Message: err.Translate(transl),
			})
		}
		return ierrors.NewErrorWithCauses(errorCauses)
	} else {
		return ierrors.NewBadRequestError()
	}
}
