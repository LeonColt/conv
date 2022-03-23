package common

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const iSO8601DateRegexString = "^(?:[1-9]\\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\\d|2[0-3]):[0-5]\\d:[0-5]\\d(?:\\.\\d{1,9})?(?:Z|[+-][01]\\d:[0-5]\\d)$"
const sha256RegexString = "^[A-Fa-f0-9]{64}$"

func NewValidator() *Validator {
	val := validator.New()
	val.RegisterValidation("ISO8601", IsISO8601Field)
	val.RegisterValidation("SHA256", IsSHA256Field)
	return &Validator{
		validator: val,
	}
}

type Validator struct {
	validator *validator.Validate
}

func IsISO8601Field(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(iSO8601DateRegexString)
	return regex.MatchString(fl.Field().String())
}

func IsSHA256Field(fl validator.FieldLevel) bool {
	return IsSHA256(fl.Field().String())
}

func IsSHA256(hash string) bool {
	regex := regexp.MustCompile(sha256RegexString)
	return regex.MatchString(hash)
}

func (validator *Validator) Validate(i interface{}) error {
	return validator.validator.Struct(i)
}
