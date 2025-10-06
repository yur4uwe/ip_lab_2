package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	encrypt := flag.Bool("e", false, "Mode: encrypt or decrypt")
	decrypt := flag.Bool("d", false, "Mode: decrypt or encrypt")
	key := flag.String("k", "", "Key")
	inputfile := flag.String("in", "", "Input file")
	outputFile := flag.String("out", "result.txt", "Output file")
	flag.Parse()

	if *key == "" {
		fmt.Println("Error: Key length is required. Use -k to specify the key.")
		os.Exit(1)
	}

	if *encrypt == *decrypt {
		fmt.Println("Error: Specify either -e (encrypt) or -d (decrypt), but not both.")
		os.Exit(1)
	}

	if *inputfile == "" {
		fmt.Println("Warning: The -in flag is ignored. Provide the input file as the first argument.")
		os.Exit(1)
	}

	inputData, err := os.ReadFile(*inputfile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	var output []byte

	if *encrypt {
		fmt.Println("Encryption Key:", *key)
		output = Encrypt(inputData, []byte(*key))
	} else if *decrypt {
		fmt.Println("Decryption Key:", *key)
		output = Decrypt(inputData, []byte(*key))
	}

	err = os.WriteFile(*outputFile, output, 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Encryption and decryption results written to %s\n", *outputFile)
}
