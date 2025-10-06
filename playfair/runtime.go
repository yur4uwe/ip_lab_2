package main

import (
	"strings"
	"unicode"
)

func createGrid(key string) [5][5]rune {
	key = strings.ToUpper(key)
	key = strings.ReplaceAll(key, "J", "I")
	used := make(map[rune]bool)
	grid := [5][5]rune{}
	alphabet := "ABCDEFGHIKLMNOPQRSTUVWXYZ"

	row, col := 0, 0
	for _, char := range key {
		if !used[char] && unicode.IsLetter(char) {
			grid[row][col] = char
			used[char] = true
			col++
			if col == 5 {
				col = 0
				row++
			}
		}
	}

	for _, char := range alphabet {
		if !used[char] {
			grid[row][col] = char
			used[char] = true
			col++
			if col == 5 {
				col = 0
				row++
			}
		}
	}

	return grid
}

func findPosition(grid [5][5]rune, char rune) (int, int) {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if grid[row][col] == char {
				return row, col
			}
		}
	}
	return -1, -1
}

func Encrypt(plaintext, key string) string {
	grid := createGrid(key)
	plaintext = strings.ToUpper(plaintext)
	plaintext = strings.ReplaceAll(plaintext, "J", "I")

	var digraphs []string
	for i := 0; i < len(plaintext); i++ {
		if !unicode.IsLetter(rune(plaintext[i])) {
			digraphs = append(digraphs, string(plaintext[i]))
			continue
		}

		if i+1 < len(plaintext) && unicode.IsLetter(rune(plaintext[i+1])) && plaintext[i] != plaintext[i+1] {
			digraphs = append(digraphs, plaintext[i:i+2])
			i++
		} else {
			digraphs = append(digraphs, string(plaintext[i])+"X")
		}
	}

	var ciphertext strings.Builder
	for _, digraph := range digraphs {
		if len(digraph) == 1 {
			ciphertext.WriteString(digraph)
			continue
		}

		row1, col1 := findPosition(grid, rune(digraph[0]))
		row2, col2 := findPosition(grid, rune(digraph[1]))

		if row1 == row2 {
			ciphertext.WriteRune(grid[row1][(col1+1)%5])
			ciphertext.WriteRune(grid[row2][(col2+1)%5])
		} else if col1 == col2 {
			ciphertext.WriteRune(grid[(row1+1)%5][col1])
			ciphertext.WriteRune(grid[(row2+1)%5][col2])
		} else {
			ciphertext.WriteRune(grid[row1][col2])
			ciphertext.WriteRune(grid[row2][col1])
		}
	}

	return ciphertext.String()
}

func Decrypt(ciphertext, key string) string {
	grid := createGrid(key)
	ciphertext = strings.ToUpper(ciphertext)

	var digraphs []string
	for i := 0; i < len(ciphertext); i++ {
		if !unicode.IsLetter(rune(ciphertext[i])) {
			digraphs = append(digraphs, string(ciphertext[i]))
			continue
		}

		if i+1 < len(ciphertext) && unicode.IsLetter(rune(ciphertext[i+1])) {
			digraphs = append(digraphs, ciphertext[i:i+2])
			i++
		} else {
			digraphs = append(digraphs, string(ciphertext[i])+"X")
		}
	}

	var plaintext strings.Builder
	for _, digraph := range digraphs {
		if len(digraph) == 1 {
			plaintext.WriteString(digraph)
			continue
		}

		row1, col1 := findPosition(grid, rune(digraph[0]))
		row2, col2 := findPosition(grid, rune(digraph[1]))

		if row1 == row2 {
			plaintext.WriteRune(grid[row1][(col1+4)%5])
			plaintext.WriteRune(grid[row2][(col2+4)%5])
		} else if col1 == col2 {
			plaintext.WriteRune(grid[(row1+4)%5][col1])
			plaintext.WriteRune(grid[(row2+4)%5][col2])
		} else {
			plaintext.WriteRune(grid[row1][col2])
			plaintext.WriteRune(grid[row2][col1])
		}
	}

	return plaintext.String()
}
