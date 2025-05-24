package faker

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[seededRand.Intn(len(letters))]
	}
	return string(b)
}

func RandomEmail() string {
	user := RandomString(8)
	domain := RandomString(5)
	return strings.ToLower(user + "@" + domain + ".test")
}

func RandomSlug(prefix string) string {
	return strings.ToLower(prefix + "-" + strconv.Itoa(seededRand.Intn(100000)))
}
