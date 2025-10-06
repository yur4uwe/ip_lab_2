#!/bin/bash

# List of all algorithms
ALGORITHMS=("vigenere" "railfence" "playfair" "blowfish")

# Function to build a specific algorithm
build_algorithm() {
    local algorithm=$1
    local folder="c:\\Users\\tsus1\\OneDrive\\Документи\\Yura_Files\\info_prot\\lab_2\\$algorithm"
    local main_file="$folder\\"
    local output_binary="$algorithm.exe"

        echo "Building $algorithm..."
        go build -o "$output_binary" "$main_file"
        if [ $? -eq 0 ]; then
            echo "$algorithm built successfully: $output_binary"
        else
            echo "Failed to build $algorithm"
        fi
}

# If no arguments are provided, build all algorithms
if [ $# -eq 0 ]; then
    echo "No algorithm specified. Building all algorithms..."
    for algorithm in "${ALGORITHMS[@]}"; do
        build_algorithm "$algorithm"
    done
else
    # Build only the specified algorithms
    echo "Building specified algorithms: $@"
    for algorithm in "$@"; do
        if [[ " ${ALGORITHMS[@]} " =~ " ${algorithm} " ]]; then
            build_algorithm "$algorithm"
        else
            echo "Unknown algorithm: $algorithm"
        fi
    done
fi