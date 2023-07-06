package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	randomString := base64.URLEncoding.EncodeToString(buffer)
	randomString = sanitizeString(randomString)

	return randomString, nil
}

func sanitizeString(str string) string {
	str = str[:len(str)-2]
	sanitizedStr := ""
	for _, ch := range str {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			sanitizedStr += string(ch)
		}
	}

	return sanitizedStr
}
