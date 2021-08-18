package helper

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

const defaultCodeLength = 8

type StringGenerator struct {
	Value string
}

// NewStringGenerator creates a new StringGenerator.
func NewStringGenerator() (*StringGenerator, error) {
	generatedString, err := generateString(defaultCodeLength)
	if err != nil {
		return nil, err
	}
	return &StringGenerator{Value: generatedString}, nil
}

func generateString(length int) (string, error) {
	if length < 1 {
		return "", errors.New("invalid code length provided")
	}
	rand.Seed(time.Now().UnixNano())
	chars := []rune(
		"abcdefghijkmnpqrstuvwxyz" +
			"123456789",
	)
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String(), nil
}
