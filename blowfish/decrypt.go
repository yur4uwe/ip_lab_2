package main

import "encoding/binary"

func Decrypt(data []byte, key []byte) []byte {
	return runBlowfish(data, key, false)
}

func decryptBlock(block []byte, localP [18]uint32) []byte {
	L := uint32(0)
	R := uint32(0)

	if len(block) >= 8 {
		L = binary.BigEndian.Uint32(block[:4])
		R = binary.BigEndian.Uint32(block[4:8])
	}

	for i := 17; i > 1; i-- {
		L ^= localP[i]
		R ^= ffunc(L)
		L, R = R, L
	}
	L, R = R, L
	R ^= localP[1]
	L ^= localP[0]

	decrypted := make([]byte, len(block))

	binary.BigEndian.PutUint32(decrypted[:4], L)
	binary.BigEndian.PutUint32(decrypted[4:8], R)

	return decrypted
}
