#!/bin/bash

ALGORITHMS=("vigenere" "railfence" "playfair" "blowfish")

build_algorithm() {
    local algorithm=$1
    local folder="c:\\Users\\tsus1\\OneDrive\\Документи\\Yura_Files\\info_prot\\lab_2\\$algorithm"
    local output_binary="$algorithm.exe"

    echo "Building $algorithm..."
    go build -o "$output_binary" "$folder"
    if [ $? -eq 0 ]; then
        echo "$algorithm built successfully: $output_binary"
    else
        echo "Failed to build $algorithm"
    fi
}

if [ $# -eq 0 ]; then
    echo "No algorithm specified. Building all algorithms..."
    for algorithm in "${ALGORITHMS[@]}"; do
        build_algorithm "$algorithm"
    done
else
    echo "Building specified algorithms: $@"
    for algorithm in "$@"; do
        if [[ " ${ALGORITHMS[@]} " =~ " ${algorithm} " ]]; then
            build_algorithm "$algorithm"
        else
            echo "Unknown algorithm: $algorithm"
        fi
    done
fi