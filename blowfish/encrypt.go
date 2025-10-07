package main

import (
	"encoding/binary"
)

func encryptBlock(block []byte, localP [18]uint32, localS [4][256]uint32) []byte {
	L := uint32(0)
	R := uint32(0)

	if len(block) >= 8 {
		L = binary.BigEndian.Uint32(block[:4])
		R = binary.BigEndian.Uint32(block[4:8])
	}

	L, R = feistelNetwork(L, R, localP, localS)

	binary.BigEndian.PutUint32(block[:4], L)
	binary.BigEndian.PutUint32(block[4:8], R)

	return block
}

func feistelNetwork(L, R uint32, localP [18]uint32, localS [4][256]uint32) (uint32, uint32) {
	for i := range 16 {
		L ^= localP[i]
		R ^= ffunc(L, localS)
		L, R = R, L
	}
	L, R = R, L
	R ^= localP[16]
	L ^= localP[17]
	return L, R
}
