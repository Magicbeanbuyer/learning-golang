package util // Package util has the same name as the directory
import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "qwertasdfgzxcvbyuiophjklnm"

// special function, init automatically runs when a package is used
func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // return a value between [min, max]
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := rand.Intn(k)
		var oneByte byte = alphabets[c]
		sb.WriteByte(oneByte)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomAmount() int64 {
	return RandomInt(20, 99999)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "RMB", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
