package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qwertyuiopasdfghjklzxcvbnm"

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + r.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency
func RandomCurrency() string {
	currencies := []string{EUR, USD, RUB}

	n := len(currencies)

	return currencies[rand.Intn(n)]
}

// RandomEmail generate a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
