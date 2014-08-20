package util

func FixedXOR(a []byte, b []byte) []byte {
	result := make([]byte, len(a))

	for index, _ := range a {
		result[index] = a[index] ^ b[index]
	}

	return result
}
