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

	localP, localS := initSAndP(key)

	for i := 0; i < blocks; i++ {
		start := i * 8
		end := start + 8
		if encrypt {
			copy(output[start:end], encryptBlock(data[start:end], localP, localS))
		} else {
			copy(output[start:end], decryptBlock(data[start:end], localP, localS))
		}
	}

	if !encrypt {
		output = removePadding(output)
	}

	return output
}

func ffunc(x uint32, Svals [4][256]uint32) uint32 {
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

func initSAndP(key []byte) ([18]uint32, [4][256]uint32) {
	localP := P
	localS := S

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
