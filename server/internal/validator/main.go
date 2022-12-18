package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"regexp"
)

var Validate = validator.New()
var usernameRegex = regexp.MustCompile("^[a-z][a-z0-9_\\-]*$")

func init() {
	err := Validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		return usernameRegex.MatchString(fl.Field().String())
	})
	if err != nil {
		logrus.Fatalln(err)
	}
}
