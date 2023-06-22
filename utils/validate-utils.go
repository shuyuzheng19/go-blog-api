package utils

import (
	"net/mail"
	"regexp"
)

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsImageURL(url string) bool {
	imageRegex := `^https?://.*\.(png|jpe?g|gif|svg|ico)$`

	regex := regexp.MustCompile(imageRegex)

	return regex.MatchString(url)
}
