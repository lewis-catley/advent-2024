package utils

import (
	"os"
)

type IBytesParser interface {
	Parse(b []byte) error
}

func ReadFile(filename string, out IBytesParser) error {
	b, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Parse the bytes onto a useable struct
	return out.Parse(b)
}
