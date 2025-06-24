package pkg_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"pt-xyz-multifinance/pkg"
)

func TestGenerateAndValidateToken_UsesRuntimeSecret(t *testing.T) {
	os.Setenv("JWT_SECRET", "secret1")
	token, err := pkg.GenerateToken("user123")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// change secret after token generation
	os.Setenv("JWT_SECRET", "secret2")
	_, err = pkg.ValidateToken(token)
	assert.Error(t, err)
}
