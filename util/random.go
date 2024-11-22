package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Randomly generate a integer between min and max
func RandomInt(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

// Randomly generate a string with length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(letter)

	for i := 0; i < n; i++ {
		c := letter[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Randomly generate a owner name
func RandomOwner() string {
	return RandomString(6)
}

// Randomly generate a currency
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Randomly generate a currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// randomly generate an email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
