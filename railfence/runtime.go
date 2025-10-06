package main

import (
	"fmt"
	"strings"
)

func Encrypt(plaintext string, rails int) string {
	direction := 1
	current_rail := 0

	fmt.Println("Rails:", rails)

	rail := make([][]rune, rails)
	for _, r := range plaintext {
		fmt.Println("Char:", string(r), "Current rail:", current_rail)
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

		fmt.Println("Rail", i, ":", string(row))
	}

	return string(ciphertext)
}

func Decrypt(ciphertext string, rails int) string {
	railLengths := make([]int, rails)
	direction := 1
	current_rail := 0

	for i := 0; i < len(ciphertext); i++ {
		railLengths[current_rail]++
		switch current_rail {
		case 0:
			direction = 1
		case rails - 1:
			direction = -1
		}
		current_rail += direction
	}

	rail := make([][]rune, rails)
	index := 0
	for i := 0; i < rails; i++ {
		rail[i] = []rune(ciphertext[index : index+railLengths[i]])
		index += railLengths[i]
	}

	plaintext := strings.Builder{}
	current_rail = 0
	direction = 1

	for i := 0; i < len(ciphertext); i++ {
		plaintext.WriteRune(rail[current_rail][0])
		rail[current_rail] = rail[current_rail][1:]

		switch current_rail {
		case 0:
			direction = 1
		case rails - 1:
			direction = -1
		}
		current_rail += direction
	}

	return plaintext.String()
}
