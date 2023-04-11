package main

import (
	"fmt"
	"strings"
)

// fungsi untuk membuat matriks Playfair Cipher dari kunci
func createMatrix(key string) [][]rune {
	key = strings.ToUpper(key)
	key = strings.ReplaceAll(key, "J", "I")
	keyMatrix := make([][]rune, 5)
	usedChars := make(map[rune]bool)

	for i := range keyMatrix {
		keyMatrix[i] = make([]rune, 5)
	}

	row, col := 0, 0
	for _, c := range key {
		if !usedChars[c] {
			keyMatrix[row][col] = c
			usedChars[c] = true
			col++
			if col == 5 {
				col = 0
				row++
			}
		}
	}

	for c := 'A'; c <= 'Z'; c++ {
		if c == 'J' {
			continue
		}
		if !usedChars[c] {
			keyMatrix[row][col] = c
			usedChars[c] = true
			col++
			if col == 5 {
				col = 0
				row++
			}
		}
	}
	return keyMatrix
}

// fungsi untuk mencari lokasi huruf dalam matriks Playfair Cipher
func findLocation(matrix [][]rune, c rune) (int, int) {
	for i, row := range matrix {
		for j, v := range row {
			if v == c {
				return i, j
			}
		}
	}
	return -1, -1
}

// enkripsi pesan menggunakan Playfair Cipher
func encrypt(plaintext, key string) string {
	matrix := createMatrix(key)

	plaintext = strings.ToUpper(plaintext)
	plaintext = strings.ReplaceAll(plaintext, "J", "I")
	ciphertext := ""

	for i := 0; i < len(plaintext); i += 2 {

		c1 := plaintext[i]
		c2 := byte('X')
		if i+1 < len(plaintext) {
			c2 = plaintext[i+1]
		}
		if c1 == c2 {
			c2 = 'X'
			i--
		}
		row1, col1 := findLocation(matrix, rune(c1))
		row2, col2 := findLocation(matrix, rune(c2))
		if row1 == row2 {
			ciphertext += string(matrix[row1][(col1+1)%5]) + string(matrix[row2][(col2+1)%5])
		} else if col1 == col2 {
			ciphertext += string(matrix[(row1+1)%5][col1]) + string(matrix[(row2+1)%5][col2])
		} else {
			ciphertext += string(matrix[row1][col2]) + string(matrix[row2][col1])
		}
	}
	return ciphertext
}

// fungsi untuk dekripsi pesan yang dienkripsi menggunakan Playfair Cipher
func decrypt(ciphertext, key string) string {
	ciphertext = strings.ToUpper(ciphertext)
	ciphertext = strings.ReplaceAll(ciphertext, "J", "I")
	matrix := createMatrix(key)
	plaintext := ""
	for i := 0; i < len(ciphertext); i += 2 {
		c1 := ciphertext[i]
		c2 := byte('X')
		if i+1 < len(ciphertext) {
			c2 = ciphertext[i+1]
		}
		row1, col1 := findLocation(matrix, rune(c1))
		row2, col2 := findLocation(matrix, rune(c2))
		if row1 == row2 {
			plaintext += string(matrix[row1][(col1+4)%5])
			plaintext += string(matrix[row2][(col2+4)%5])
		} else if col1 == col2 {
			plaintext += string(matrix[(row1+4)%5][col1])
			plaintext += string(matrix[(row2+4)%5][col2])
		} else {
			plaintext += string(matrix[row1][col2])
			plaintext += string(matrix[row2][col1])
		}
	}
	return plaintext
}

func main() {
	plaintext := "tafanianatalia"
	key := "untad"

	fmt.Println("Plaintext: ", plaintext)
	fmt.Println("Key: ", key)

	ciphertext := encrypt(plaintext, key)
	fmt.Println("Ciphertext: ", ciphertext)

	decryptedPlaintext := decrypt(ciphertext, key)
	fmt.Println("Decrypted plaintext: ", decryptedPlaintext)
}
