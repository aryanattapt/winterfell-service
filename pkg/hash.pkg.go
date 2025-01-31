package pkg

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncodeBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func DecodeBase64(data string) string {
	var decodedByte, _ = base64.StdEncoding.DecodeString(data)
	return string(decodedByte)
}

func HashSHA1(data string) string {
	var hash = sha1.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func HashSHA256(data string) string {
	var hash = sha256.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func HashMD5(data string) string {
	var hash = md5.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func HashPasswordBCrypt(data string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	return string(hash)
}

func ComparePasswordBCrypt(data string, hashedData string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedData), []byte(data)); err != nil {
		return false
	} else {
		return true
	}
}
