package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// init set rand seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

// RandomString generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(len(alphabet))]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generate a random owner
func RandomOwner() string {
	return RandomString(8)
}

// RandomMoney generate a random money
func RandomMoney() int64 {
	return RandomInt(100, 500)
}

// RandomCurrency generate a random currency
func RandomCurrency() string {
	currencies := []string{"RMB", "EUR", "USD"}
	return currencies[rand.Intn(len(currencies))]
}

// RandomEmail generate a random email
func RandomEmail() string {
	email := fmt.Sprintf("%s@email.com", RandomString(6))
	return email
}
