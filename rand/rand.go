package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// Bytes will generate n random bytes or return an error if there was one. This
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	nRead, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("bytes : %w", err)
	}
	if nRead < n {
		return nil, fmt.Errorf("bytes : didnt read enough random bytes")
	}
	return b, nil
}

// String will generate a random string of n bytes or return an error if there was one. This
func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("string : %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

const SessionTokenBytes = 32

// SessionToken is a helper function designed to generate
// session tokens of a predetermined byte size.
func SessionToken() (string, error) {
	return String(SessionTokenBytes)
}
