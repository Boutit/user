package utils

import (
	"net/mail"
)


func IsEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

func IsPhoneNumber(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}