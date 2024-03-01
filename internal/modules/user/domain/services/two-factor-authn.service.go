package services

import (
	"context"
	"fmt"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/security"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/pquerna/otp/totp"
)

func (u *UserService) GenerateTOTP(id string) (map[string]string, error) {

	user, err := u.FindById(id, context.Background())
	if err != nil {
		return nil, errors.ErrCtx(err, "u.FindById")
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "financial-file-manager.com",
		AccountName: user.Email,
		SecretSize:  15,
	})
	if err != nil {
		return nil, errors.ErrCtx(err, "totp.Generate")
	}

	err = u.encryptTOTP(user, key.Secret(), key.URL())
	if err != nil {
		return nil, errors.ErrCtx(err, "u.encryptTOTP")
	}

	err = u.Repository.UpdateOTP(user, context.Background())
	if err != nil {
		return nil, errors.ErrCtx(err, "u.Repository.UpdateOTP")
	}

	otpResponse := map[string]string{
		"base32":      key.Secret(),
		"otpauth_url": key.URL(),
	}
	return otpResponse, nil
}

func (u *UserService) VerifyTOTP(id, token, isValidate string) error {
	aes_key := config.GetEnv("security_aes_key").(string)

	user, err := u.FindById(id, context.Background())
	if err != nil {
		return errors.ErrCtx(err, "u.FindById")
	}

	otpSecret, err := security.Decrypt(user.OtpSecret, aes_key)
	if err != nil {
		return errors.ErrCtx(err, "security.Decrypt OtpSecret")
	}	

	valid := totp.Validate(token, otpSecret)
	if !valid {
		return errors.ErrCtx(fmt.Errorf("token is invalid"), "u.FindById")
	}

	if isValidate == "1" {		
		user.OtpEnabled = true
	}

	err = u.Repository.UpdateOTPVerify(user, context.Background())
	if err != nil {
		return errors.ErrCtx(err, "u.Repository.UpdateOTPVerify")
	}
	return nil
}

func (u *UserService) DisableOTP(id string) error {
	user, err := u.FindById(id, context.Background())
	if err != nil {
		return errors.ErrCtx(err, "u.FindById")
	}

	user.OtpEnabled = false

	err = u.Repository.UpdateOTPVerify(user, context.Background())
	if err != nil {
		return errors.ErrCtx(err, "u.Repository.UpdateOTPVerify")
	}
	return nil
}

func (u *UserService) encryptTOTP(user *entities.User, otpSecret, otpAuthUrl string) error {
	aes_key := config.GetEnv("security_aes_key").(string)
	var err error

	user.OtpSecret, err = security.Encrypt(otpSecret, aes_key)
	if err != nil {
		return errors.ErrCtx(err, "security.Encrypt LastName")
	}
	user.OtpAuthUrl, err = security.Encrypt(otpAuthUrl, aes_key)
	if err != nil {
		return errors.ErrCtx(err, "security.Encrypt Email")
	}
	return nil
}