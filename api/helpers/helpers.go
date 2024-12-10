package helpers

import (
	"os"
	"strings"
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
