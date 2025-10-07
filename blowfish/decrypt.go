package main

import (
	"encoding/binary"
)

func decryptBlock(block []byte, localP [18]uint32, localS [4][256]uint32) []byte {
	L := uint32(0)
	R := uint32(0)

	if len(block) >= 8 {
		L = binary.BigEndian.Uint32(block[:4])
		R = binary.BigEndian.Uint32(block[4:8])
	}

	L, R = reverseFeistelNetwork(L, R, localP, localS)

	decrypted := make([]byte, len(block))

	binary.BigEndian.PutUint32(decrypted[:4], L)
	binary.BigEndian.PutUint32(decrypted[4:8], R)

	return decrypted
}

func reverseFeistelNetwork(L, R uint32, localP [18]uint32, localS [4][256]uint32) (uint32, uint32) {
	for i := 17; i > 1; i-- {
		L ^= localP[i]
		R ^= ffunc(L, localS)
		L, R = R, L
	}
	L, R = R, L
	R ^= localP[1]
	L ^= localP[0]
	return L, R
}
