package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/security"
	"github.com/golang/mock/gomock"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTOTP(t *testing.T) {
	userService, mockRepo, _, mockMemcached := prepare(t)

	mockMemcached.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()
	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes()
	aes_key := config.GetEnv("security_aes_key").(string)

	t.Run("Should return base32 secret and otpauth_url", func(t *testing.T) {
		user, _ := entities.NewUser("John", "Doe", "123.356.229-00", "john.doe@example.com", password)

		mockRepo.EXPECT().
			FindById(user.ID, gomock.Any()).
			Return(user, nil).Times(1)

		encryptedSecret, _ := security.Encrypt(user.OtpSecret, aes_key)
		encryptedURL, _ := security.Encrypt(user.OtpAuthUrl, aes_key)
		encryptedEmail, _ := security.Encrypt(user.Email, aes_key)
		encryptedCPF, _ := security.Encrypt(user.CPF, aes_key)
		encryptedLastName, _ := security.Encrypt(user.LastName, aes_key)

		user.Email = encryptedEmail
		user.CPF = encryptedCPF
		user.LastName = encryptedLastName
		user.OtpSecret = encryptedSecret
		user.OtpAuthUrl = encryptedURL

		mockRepo.EXPECT().UpdateOTP(user, gomock.Any()).Return(nil).Times(1)

		response, err := userService.GenerateTOTP(user.ID.String())
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Contains(t, response, "base32")
		assert.Contains(t, response, "otpauth_url")
	})

}

func TestVerifyTOTP(t *testing.T) {

	user, _ := entities.NewUser("John", "Doe", "123.356.229-00", "john.doe@example.com", password)

	userService, mockRepo, _, mockMemcached := prepare(t)
	aes_key := config.GetEnv("security_aes_key").(string)

	mockMemcached.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()
	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes()
	t.Run("Should return nil when successfully 2FA valid authn", func(t *testing.T) {
		mockRepo.EXPECT().
			FindById(user.ID, gomock.Any()).
			Return(user, nil).AnyTimes()

		user.OtpSecret = "JKHASD675546SAD6"
		encryptedSecret, _ := security.Encrypt(user.OtpSecret, aes_key)
		encryptedURL, _ := security.Encrypt(user.OtpAuthUrl, aes_key)
		encryptedEmail, _ := security.Encrypt(user.Email, aes_key)
		encryptedCPF, _ := security.Encrypt(user.CPF, aes_key)
		encryptedLastName, _ := security.Encrypt(user.LastName, aes_key)

		user.Email = encryptedEmail
		user.CPF = encryptedCPF
		user.LastName = encryptedLastName
		user.OtpSecret = encryptedSecret
		user.OtpAuthUrl = encryptedURL

		otpSecret, _ := security.Decrypt(user.OtpSecret, aes_key)

		user.OtpEnabled = true

		token, _ := totp.GenerateCode(otpSecret, time.Now())

		mockRepo.EXPECT().UpdateOTPVerify(user, gomock.Any()).Return(nil).Times(1)

		err := userService.VerifyTOTP(user.ID.String(), token, "1")
		assert.NoError(t, err)
	})
}

func TestDisableOTP(t *testing.T) {
	user, _ := entities.NewUser("John", "Doe", "123.356.229-00", "john.doe@example.com", password)
	aes_key := config.GetEnv("security_aes_key").(string)

	userService, mockRepo, _, mockMemcached := prepare(t)

	mockMemcached.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()
	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes()
	t.Run("Should return nil when successfully 2FA valid authn", func(t *testing.T) {
		mockRepo.EXPECT().
			FindById(user.ID, gomock.Any()).
			Return(user, nil).AnyTimes()

		encryptedSecret, _ := security.Encrypt(user.OtpSecret, aes_key)
		encryptedURL, _ := security.Encrypt(user.OtpAuthUrl, aes_key)
		encryptedEmail, _ := security.Encrypt(user.Email, aes_key)
		encryptedCPF, _ := security.Encrypt(user.CPF, aes_key)
		encryptedLastName, _ := security.Encrypt(user.LastName, aes_key)

		user.Email = encryptedEmail
		user.CPF = encryptedCPF
		user.LastName = encryptedLastName
		user.OtpSecret = encryptedSecret
		user.OtpAuthUrl = encryptedURL

		user.OtpEnabled = false

		mockRepo.EXPECT().UpdateOTPVerify(user, context.Background()).
			Return(nil).Times(1)

		err := userService.DisableOTP(user.ID.String())
		assert.NoError(t, err)
	})

}
