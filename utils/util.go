package utils

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

func IsEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
func IsPhoneValid(phone string) bool {
	phoneRegex := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	return phoneRegex.MatchString(strings.TrimSpace(phone))
}

func IsUrlValid(URL string) bool {
	_, err := url.ParseRequestURI(URL)
	return err == nil
}

func IsValidUUID(input uuid.UUID) bool {
	_, err := uuid.Parse(input.String())
	return err == nil
}
