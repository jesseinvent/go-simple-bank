package util

import (
	"database/sql"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomSqlInt(id int64) sql.NullInt64 {
	return sql.NullInt64{Int64: RandomInt(1, 100), Valid: true}
}

// Generates a random string of n characters
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// Generates a random money name
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "NGN"}

	n := len(currencies)

	return currencies[rand.Intn((n))]
}
