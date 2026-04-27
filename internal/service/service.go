package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertAuto(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", errors.New("empty input")
	}

	if isMorse(input) {
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}

func isMorse(input string) bool {
	input = strings.TrimSpace(input)

	for _, s := range input {
		switch s {
		case '.', '-', '/', ' ', '\n', '\r', '\t':
			continue
		default:
			return false
		}
	}
	return true
}
