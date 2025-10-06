package main

import "strings"

func Encrypt(plaintext string, rails int) string {
	direction := 1
	current_rail := 0

	rail := make([][]rune, rails)
	for _, r := range plaintext {
		switch current_rail {
		case 0:
			direction = 1
		case rails - 1:
			direction = -1
		}

		rail[current_rail] = append(rail[current_rail], r)
		current_rail += direction
	}

	var ciphertext []rune
	for i, row := range rail {
		ciphertext = append(ciphertext, row...)
		if i < len(rail)-1 {
			ciphertext = append(ciphertext, ' ')
		}
	}

	return string(ciphertext)
}

func Decrypt(ciphertext string, rails int) string {
	words := strings.Split(ciphertext, " ")
	if len(words) != rails {
		return string("")
	}

	direction := 1
	current_rail := 0
	plaintext := strings.Builder{}

	for i := 0; i < len(ciphertext); i++ {
		switch current_rail {
		case 0:
			direction = 1
		case rails - 1:
			direction = -1
		}

		if len(words[current_rail]) > 0 {
			plaintext.WriteByte(words[current_rail][0])
			words[current_rail] = words[current_rail][1:]
		}
		current_rail += direction
	}

	return plaintext.String()
}
