package validate

import (
	go_err "errors"
	errors "github.com/Lucasvmarangoni/logella/err"
	"regexp"
)

func ValidatePassword(password string) error {
	if len(password) < 10 {
		return errors.ErrCtx(go_err.New("password must be at least 10 characters long"), "validatePassword")
	}

	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.ErrCtx(go_err.New("password must contain at least one lowercase letter"), "validatePassword")
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.ErrCtx(go_err.New("password must contain at least one uppercase letter"), "validatePassword")
	}

	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return errors.ErrCtx(go_err.New("password must contain at least one digit"), "validatePassword")
	}

	if !regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(password) {
		return errors.ErrCtx(go_err.New("password must contain at least one special character"), "validatePassword")
	}
	return nil
}
