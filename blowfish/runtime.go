package main

import (
	"encoding/binary"
	"io"
	"os"
)

func processBlock(block []byte, localP p, localS s, encrypt bool) []byte {
	if encrypt {
		block = applyPadding(block, blockSize)
		return encryptBlock(block, localP, localS)
	} else {
		block = decryptBlock(block, localP, localS)
		return removePadding(block)
	}
}

func Stream(infile *os.File, outfile *os.File, key []byte, encrypt bool) error {

	localP, localS := initializeBlowfishKey(key)

	buffer := make([]byte, blockSize)

	for {
		n, err := infile.Read(buffer)
		if err != nil && err != io.EOF {
			break
		}

		if n == 0 {
			break
		}

		block := processBlock(buffer[:n], localP, localS, encrypt)

		_, write_err := outfile.Write(block)
		if write_err != nil {
			return write_err
		}

		if err == io.EOF {
			break
		}
	}

	return nil
}

func ffunc(x uint32, Svals s) uint32 {
	byte0 := byte(x >> 24)
	byte1 := byte((x >> 16) & 0xFF)
	byte2 := byte((x >> 8) & 0xFF)
	byte3 := byte(x & 0xFF)

	res := Svals[0][byte0]
	res = res + Svals[1][byte1]
	res = res ^ Svals[2][byte2]
	res = res + Svals[3][byte3]

	return res
}

func initializeBlowfishKey(key []byte) (p, s) {
	localP := P
	localS := S

	key = applyPadding(key, 4)

	keyLen := len(key)
	j := 0
	for i := 0; i < 18; i++ {
		localP[i] ^= binary.BigEndian.Uint32(key[j%keyLen : (j%keyLen)+4])
		j += 4
	}

	L, R := uint32(0), uint32(0)
	for i := 0; i < 18; i += 2 {
		L, R = feistelNetwork(L, R, localP, localS)
		localP[i], localP[i+1] = L, R
	}

	for s := 0; s < 4; s++ {
		for i := 0; i < 256; i += 2 {
			L, R = feistelNetwork(L, R, localP, localS)
			localS[s][i], localS[s][i+1] = L, R
		}
	}

	return localP, localS
}

func applyPadding(data []byte, blockSize int) []byte {
	if len(data)%blockSize == 0 {
		return data
	}

	padding := blockSize - (len(data) % blockSize)
	if padding == 0 {
		padding = blockSize
	}

	padded := make([]byte, len(data)+padding)
	copy(padded, data)

	for i := len(data); i < len(padded); i++ {
		padded[i] = byte(padding)
	}

	return padded
}

func removePadding(data []byte) []byte {
	if len(data) == 0 {
		return data
	}

	padding := int(data[len(data)-1])
	if padding > len(data) || padding == 0 {
		return data
	}

	return data[:len(data)-padding]
}
