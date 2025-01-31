package pkg

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func IsEmptyString(data string) bool {
	return len(strings.TrimSpace(data)) == 0 || strings.TrimSpace(data) == ""
}

func GenerateRandomString(length int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func GenerateRandomNumber(length int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var letters = []rune("1234567890")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func GenerateUUID() string {
	return uuid.New().String()
}

func LowercaseFirstChar(s string) string {
	if len(s) == 0 {
		return s
	}
	firstChar := strings.ToLower(string(s[0]))
	return firstChar + s[1:]
}
