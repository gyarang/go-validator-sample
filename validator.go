package main

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

type SignUp struct {
	ID       string `validate:"required,gte=4,lte=255,alphanum"`
	Password string `validate:"required,password"`
	Gender   string `validate:"gender"`
	Age      int    `validate:"gte=0"`
	Email    string `validate:"email"`
}

func NewSignUpValidator() *validator.Validate {
	v := validator.New()
	_ = v.RegisterValidation("password", isValidPassword)
	_ = v.RegisterValidation("gender", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "male" || fl.Field().String() == "female"
	})
	return v
}

/*
* Password rules:
* at least 7 letters
* at least 1 number
* at least 1 upper case
* at least 1 special character
 */
func isValidPassword(fl validator.FieldLevel) bool {
	field := fl.Field()
	input := field.String()

	if len(input) < 7 {
		return false
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, c := range input {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}
