package security_test

import (
	"testing"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/security"
	"github.com/stretchr/testify/require"
)

const isBase64 = "^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$"

func TestEncryptDecryptMessage(t *testing.T) {
	key := "0123456789abcdef"
	message := "Lorem ipsum dolor sit amet"

	encrypted, err := security.Encrypt(message, key)
	require.Nil(t, err)
	require.Regexp(t, isBase64, encrypted)

	decrypted, err := security.Decrypt(encrypted, key)
	require.Nil(t, err)
	require.Equal(t, message, decrypted)
}
