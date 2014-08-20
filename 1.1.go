package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	encoded_hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected_b64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	bytes, _ := hex.DecodeString(encoded_hex)

	generated_b64 := base64.StdEncoding.EncodeToString(bytes)

	if expected_b64 == generated_b64 {
		fmt.Println("Expected and generated Base64 Match")
	}
}
