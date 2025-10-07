package main

import (
	"flag"
	"fmt"
	"os"
)

func Encrypt(in, out *os.File, key []byte) error {
	return ConcurrentStream(in, out, key, true)
}

func Decrypt(in, out *os.File, key []byte) error {
	return ConcurrentStream(in, out, key, false)
}

func main() {
	encrypt := flag.Bool("e", false, "Mode: encrypt or decrypt")
	decrypt := flag.Bool("d", false, "Mode: decrypt or encrypt")
	key := flag.String("k", "", "Key")
	input_file_name := flag.String("in", "", "Input file")
	output_file_name := flag.String("out", "result.txt", "Output file")
	flag.Parse()

	if *key == "" {
		fmt.Println("Error: Key length is required. Use -k to specify the key.")
		os.Exit(1)
	}

	if *encrypt == *decrypt {
		fmt.Println("Error: Specify either -e (encrypt) or -d (decrypt), but not both.")
		os.Exit(1)
	}

	if *input_file_name == "" {
		fmt.Println("Warning: The -in flag is ignored. Provide the input file as the first argument.")
		os.Exit(1)
	}

	inFile, err := os.Open(*input_file_name)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}
	defer inFile.Close()

	outFile, err := os.OpenFile(*output_file_name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error opening/creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	if *encrypt {
		fmt.Println("Encryption Key:", *key)
		Encrypt(inFile, outFile, []byte(*key))
	} else if *decrypt {
		fmt.Println("Decryption Key:", *key)
		Decrypt(inFile, outFile, []byte(*key))
	}

	if *encrypt {
		fmt.Printf("Encryption results written to %s\n", *output_file_name)
	} else if *decrypt {
		fmt.Printf("Decryption results written to %s\n", *output_file_name)
	}
}
