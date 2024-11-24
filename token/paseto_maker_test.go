package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yantong1775/vaultCore/util"
)

func TestPasetoMaker(t *testing.T) {
	// Create a new PasetoMaker
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	// Create a new token
	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Verify the token
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	// Create a new PasetoMaker
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	// Create a new token
	username := util.RandomOwner()
	duration := -time.Minute

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Verify the token
	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestPasetoInvalidKey(t *testing.T) {
	// Create a new PasetoMaker
	_, err := NewPasetoMaker(util.RandomString(1))
	require.Error(t, err)
	require.EqualError(t, err, "invalid key size: must be exactly 32 characters")
}
