package validation

import (
	"errors"
	"unicode"
)

func ValidatePassword(password string) error {
	var hasMinLen bool = len(password) >= 8
	var hasNumber, hasUpper, hasLower, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasMinLen {
		return errors.New("şifre en az 8 karakter olmalıdır")
	}
	if !hasNumber {
		return errors.New("şifre en az bir rakam içermelidir")
	}
	if !hasUpper {
		return errors.New("şifre en az bir büyük harf içermelidir")
	}
	if !hasLower {
		return errors.New("şifre en az bir küçük harf içermelidir")
	}
	if !hasSpecial {
		return errors.New("şifre en az bir özel karakter içermelidir")
	}

	return nil
}
