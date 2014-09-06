package xor

func Fixed(a []byte, b []byte) []byte {
	result := make([]byte, len(a))

	for index, _ := range a {
		result[index] = a[index] ^ b[index]
	}

	return result
}

func SingleByte(input []byte, key byte) []byte {
	key_bytes := make([]byte, len(input))
	for i, _ := range key_bytes {
		key_bytes[i] = key
	}

	return Fixed(input, key_bytes)
}
