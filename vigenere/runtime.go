package vigenere

import (
	"strings"
)

func run(input, key string, mode bool) string {
	if len(key) != len(input) {
		key = strings.Repeat(key, len(input)/len(key)+1)[:len(input)]
	}

	var ciphertext strings.Builder
	for i := 0; i < len(input); i++ {
		t := input[i]
		k := key[i]

		if mode {
			ciphertext.WriteByte((t + k) % 255)
		} else {
			ciphertext.WriteByte((t - k) % 255)
		}
	}

	return ciphertext.String()
}

func Encrypt(plaintext, key string) string {
	return run(plaintext, key, true)
}

func Decrypt(ciphertext, key string) string {
	return run(ciphertext, key, false)
}
