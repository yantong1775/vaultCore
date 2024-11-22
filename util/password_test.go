package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	// Correct password
	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)
	// Wrong password
	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPasswordDup, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPasswordDup)
	require.NotEqual(t, hashedPassword, hashedPasswordDup)
}
