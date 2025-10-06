package main

import "encoding/binary"

func runBlowfish(data, key []byte, encrypt bool) []byte {
	if len(key) < 4 || len(key) > 72 {
		panic("Key length must be between 4 and 72 bytes")
	}

	key = applyPadding(key, 4)
	if encrypt {
		data = applyPadding(data, 8)
	}

	output := make([]byte, len(data))
	blocks := len(data) / 8

	localP := initP(key)

	for i := 0; i < blocks; i++ {
		start := i * 8
		end := start + 8
		if encrypt {
			copy(output[start:end], encryptBlock(data[start:end], localP))
		} else {
			copy(output[start:end], decryptBlock(data[start:end], localP))
		}
	}

	if !encrypt {
		output = removePadding(output)
	}

	return output
}

func ffunc(x uint32) uint32 {
	byte0 := byte(x >> 24)
	byte1 := byte((x >> 16) & 0xFF)
	byte2 := byte((x >> 8) & 0xFF)
	byte3 := byte(x & 0xFF)

	res := S[0][byte0]
	res = res + S[1][byte1]
	res = res ^ S[2][byte2]
	res = res + S[3][byte3]

	return res
}

func initP(key []byte) [18]uint32 {
	localP := P

	keyLen := len(key)
	j := 0
	for i := 0; i < 18; i++ {
		localP[i] ^= binary.BigEndian.Uint32(key[j%keyLen : (j%keyLen)+4])
		j += 4
	}

	return localP
}

func applyPadding(data []byte, blockSize int) []byte {
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
		panic("Data is empty, cannot remove padding")
	}

	padding := int(data[len(data)-1])
	if padding > len(data) || padding == 0 {
		panic("Invalid padding")
	}

	return data[:len(data)-padding]
}
