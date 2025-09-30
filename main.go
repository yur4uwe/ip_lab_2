package main

import (
	"fmt"
	"lab/blowfish"
	"lab/playfair"
	"lab/railfence"
	"lab/vigenere"
)

func main() {
	enc_target := []byte("Hello,World!")
	key := []byte("mysecretkey1234")

	enc_blowfish := blowfish.Encrypt(enc_target, key)
	dec_blowfish := blowfish.Decrypt(enc_blowfish, key)

	fmt.Println("Original:", string(enc_target))
	fmt.Println("Key:", string(key))

	fmt.Println("\nBlowfish Encryption/Decryption")
	fmt.Println("Encrypted:", string(enc_blowfish))
	fmt.Println("Decrypted:", string(dec_blowfish))

	enc_vigenere := vigenere.Encrypt(string(enc_target), string(key))
	dec_vigenere := vigenere.Decrypt(enc_vigenere, string(key))

	fmt.Println("\nVigenere Encryption/Decryption")
	fmt.Println("Encrypted:", string(enc_vigenere))
	fmt.Println("Decrypted:", string(dec_vigenere))

	enc_railfence := railfence.Encrypt(string(enc_target), 3)
	dec_railfence := railfence.Decrypt(enc_railfence, 3)

	fmt.Println("\nRailfence Encryption/Decryption")
	fmt.Println("Encrypted:", string(enc_railfence))
	fmt.Println("Decrypted:", string(dec_railfence))

	enc_playfair := playfair.Encrypt(string(enc_target), string(key))
	dec_playfair := playfair.Decrypt(enc_playfair, string(key))

	fmt.Println("\nPlayfair Encryption/Decryption")
	fmt.Println("Encrypted:", string(enc_playfair))
	fmt.Println("Decrypted:", string(dec_playfair))
}
