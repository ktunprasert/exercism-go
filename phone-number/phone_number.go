package phonenumber

import (
	"errors"
	"regexp"
)

func Number(phoneNumber string) (string, error) {
	phoneNumber = regexp.MustCompile(`\D`).ReplaceAllString(phoneNumber, "")

	switch {
	case len(phoneNumber) < 10, len(phoneNumber) > 11, len(phoneNumber) == 11 && phoneNumber[0] != '1':
		return "", errors.New("invalid phone number")
	case len(phoneNumber) == 11:
		phoneNumber = phoneNumber[1:]
	}

	if rune(phoneNumber[0]) < '2' || rune(phoneNumber[3]) < '2' {
		return "", errors.New("invalid phone number")
	}

	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return phoneNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return "(" + phoneNumber[:3] + ") " + phoneNumber[3:6] + "-" + phoneNumber[6:], nil
}
