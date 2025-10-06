package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	encrypt := flag.Bool("e", false, "Mode: encrypt or decrypt")
	decrypt := flag.Bool("d", false, "Mode: decrypt or encrypt")
	key := flag.Int("k", 0, "Key")
	inputfile := flag.String("in", "", "Input file")
	outputFile := flag.String("out", "result.txt", "Output file")
	flag.Parse()

	if *key == 0 {
		fmt.Println("Error: Key is required. Use -k to specify the key.")
		os.Exit(1)
	}

	if *encrypt == *decrypt {
		fmt.Println("Error: Specify either -e (encrypt) or -d (decrypt), but not both.")
		os.Exit(1)
	}

	if *inputfile == "" {
		fmt.Println("Error: Input file is required. Use -in to specify the input file.")
		os.Exit(1)
	}

	inputData, err := os.ReadFile(*inputfile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	const placeholder = "{{SPACE}}"

	var output string

	if *encrypt {
		plaintext := strings.ReplaceAll(string(inputData), " ", placeholder)
		output = Encrypt(plaintext, *key)
	} else if *decrypt {
		ciphertext := string(inputData)
		decrypted := Decrypt(ciphertext, *key)
		output = strings.ReplaceAll(decrypted, placeholder, " ")
	}

	err = os.WriteFile(*outputFile, []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		os.Exit(1)
	}

	if *encrypt {
		fmt.Printf("Encryption results written to %s\n", *outputFile)
	} else if *decrypt {
		fmt.Printf("Decryption results written to %s\n", *outputFile)
	}
}
