package main

import (
	"encoding/binary"
)

func Encrypt(data []byte, key []byte) []byte {
	return runBlowfish(data, key, true)
}

func encryptBlock(block []byte, localP [18]uint32) []byte {
	L := uint32(0)
	R := uint32(0)

	if len(block) >= 8 {
		L = binary.BigEndian.Uint32(block[:4])
		R = binary.BigEndian.Uint32(block[4:8])
	}

	for i := range 16 {
		L ^= localP[i]
		R ^= ffunc(L)
		L, R = R, L
	}

	L, R = R, L
	R ^= localP[16]
	L ^= localP[17]

	binary.BigEndian.PutUint32(block[:4], L)
	binary.BigEndian.PutUint32(block[4:8], R)

	return block
}
