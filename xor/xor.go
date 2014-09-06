package xor

// XORs two byte-arrrays of the same length together and returns the result
func Fixed(a []byte, b []byte) []byte {
	result := make([]byte, len(a))

	for index, _ := range a {
		result[index] = a[index] ^ b[index]
	}

	return result
}

// XORs plaintext bytes with a single key byte and returns the result
func SingleByte(plaintext []byte, key byte) []byte {
	key_bytes := make([]byte, len(plaintext))
	for i, _ := range key_bytes {
		key_bytes[i] = key
	}

	return Fixed(plaintext, key_bytes)
}

// XORs plaintext with a repeating key
func RepeatingKey(plaintext []byte, key []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	keyLength := len(key)
	for i, thisByte := range plaintext {
		ciphertext[i] = thisByte ^ key[i%keyLength]
	}
	return ciphertext
}
