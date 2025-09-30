package playfair

import (
	"strings"
	"unicode"
)

// Create a 5x5 Playfair cipher grid from the key
func createGrid(key string) [5][5]rune {
	key = strings.ToUpper(key)
	key = strings.ReplaceAll(key, "J", "I") // Treat I and J as the same letter
	used := make(map[rune]bool)
	grid := [5][5]rune{}
	alphabet := "ABCDEFGHIKLMNOPQRSTUVWXYZ"

	// Fill the grid with the key
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

	// Fill the remaining spaces with the rest of the alphabet
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

// Find the position of a character in the grid
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

// Encrypt plaintext using the Playfair cipher
func Encrypt(plaintext, key string) string {
	grid := createGrid(key)
	plaintext = strings.ToUpper(plaintext)
	plaintext = strings.ReplaceAll(plaintext, "J", "I") // Treat I and J as the same letter

	// Prepare plaintext digraphs
	var digraphs []string
	for i := 0; i < len(plaintext); i++ {
		if !unicode.IsLetter(rune(plaintext[i])) {
			// Skip non-alphabetic characters
			digraphs = append(digraphs, string(plaintext[i]))
			continue
		}

		if i+1 < len(plaintext) && unicode.IsLetter(rune(plaintext[i+1])) && plaintext[i] != plaintext[i+1] {
			digraphs = append(digraphs, plaintext[i:i+2])
			i++
		} else {
			// Add filler character 'X' if needed
			digraphs = append(digraphs, string(plaintext[i])+"X")
		}
	}

	// Encrypt each digraph
	var ciphertext strings.Builder
	for _, digraph := range digraphs {
		if len(digraph) == 1 {
			// Non-alphabetic characters are added as-is
			ciphertext.WriteString(digraph)
			continue
		}

		row1, col1 := findPosition(grid, rune(digraph[0]))
		row2, col2 := findPosition(grid, rune(digraph[1]))

		if row1 == row2 {
			// Same row: shift right
			ciphertext.WriteRune(grid[row1][(col1+1)%5])
			ciphertext.WriteRune(grid[row2][(col2+1)%5])
		} else if col1 == col2 {
			// Same column: shift down
			ciphertext.WriteRune(grid[(row1+1)%5][col1])
			ciphertext.WriteRune(grid[(row2+1)%5][col2])
		} else {
			// Rectangle rule: swap columns
			ciphertext.WriteRune(grid[row1][col2])
			ciphertext.WriteRune(grid[row2][col1])
		}
	}

	return ciphertext.String()
}

// Decrypt ciphertext using the Playfair cipher
func Decrypt(ciphertext, key string) string {
	grid := createGrid(key)
	ciphertext = strings.ToUpper(ciphertext)

	// Prepare ciphertext digraphs
	var digraphs []string
	for i := 0; i < len(ciphertext); i++ {
		if !unicode.IsLetter(rune(ciphertext[i])) {
			// Skip non-alphabetic characters
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

	// Decrypt each digraph
	var plaintext strings.Builder
	for _, digraph := range digraphs {
		if len(digraph) == 1 {
			// Non-alphabetic characters are added as-is
			plaintext.WriteString(digraph)
			continue
		}

		row1, col1 := findPosition(grid, rune(digraph[0]))
		row2, col2 := findPosition(grid, rune(digraph[1]))

		if row1 == row2 {
			// Same row: shift left
			plaintext.WriteRune(grid[row1][(col1+4)%5])
			plaintext.WriteRune(grid[row2][(col2+4)%5])
		} else if col1 == col2 {
			// Same column: shift up
			plaintext.WriteRune(grid[(row1+4)%5][col1])
			plaintext.WriteRune(grid[(row2+4)%5][col2])
		} else {
			// Rectangle rule: swap columns
			plaintext.WriteRune(grid[row1][col2])
			plaintext.WriteRune(grid[row2][col1])
		}
	}

	return plaintext.String()
}
