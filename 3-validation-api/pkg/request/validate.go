package request

import (
	"github.com/go-playground/validator"
)

func Validate[T interface{}](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}
