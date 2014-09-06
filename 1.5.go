package main

import (
	"encoding/hex"
	"fmt"
	"github.com/Najaf/cryptopals-solutions/xor"
)

func main() {
	key := []byte("ICE")
	plaintext := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	ciphertext := xor.RepeatingKey(plaintext, key)
	encoded_ciphertext := hex.EncodeToString(ciphertext)
	expected_encoded_ciphertext := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	fmt.Printf("Expected:\t%s\n", expected_encoded_ciphertext)
	fmt.Printf("Generated:\t%s\n", encoded_ciphertext)

	if expected_encoded_ciphertext == encoded_ciphertext {
		fmt.Println("Ciphertexts match")
	} else {
		fmt.Println("Ciphertexts don't match")
	}
}
