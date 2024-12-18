package helpers

import (
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func EnforeHTTP(url string) string {
	if !strings.HasPrefix(url, "http") {
		return "http://" + url
	}
	return url
}

func DetectDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return true
	}

	newUrl := strings.Replace(url, "http://", "", 1)
	newUrl = strings.Replace(url, "https://", "", 1)
	newUrl = strings.Replace(url, "www.", "", 1)
	newUrl = strings.Split(url, "/")[0]

	if newUrl == os.Getenv("DOMAIN") {
		return true
	}
	return false
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
}
