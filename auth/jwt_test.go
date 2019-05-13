package auth

import (
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"

	"github.com/hookactions/go-utils/crypto"
)

func TestNewJWTWithClaims(t *testing.T) {
	key, err := crypto.GenerateRsaKey()
	require.NoError(t, err)

	token, err := NewJWTWithClaims(jwt.MapClaims{}, key)
	require.NoError(t, err)
	require.NotEmpty(t, token)
}
